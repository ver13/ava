package hash

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	_hash "hash"
	"hash/fnv"
	"reflect"
	"sort"
	"strconv"
	"strings"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorHashAVA "github.com/ver13/ava/pkg/common/hash/error"
)

type Hash struct {
	version string
	opts    *HashOptions
}

// HashOptions are options that are available for hashing.
type HashOptions struct {
	// Hasher is the hash function to use. If this isn't set, it will default to FNV.
	Hasher _hash.Hash64

	// TagName is the struct tag to look at when hashing the structure.
	// By default this is "hash".
	TagName string

	// ZeroNil is flag determining if nil pointer should be treated equal to a zero value of pointed type. By default this is false.
	ZeroNil bool
}

func NewHash(version string) *Hash {
	opts := &HashOptions{
		Hasher:  fnv.New64(),
		TagName: "hash",
		ZeroNil: true,
	}
	// Reset the hash
	opts.Hasher.Reset()

	return &Hash{
		version: version,
		opts:    opts,
	}
}

// version returns the version of the supplied hash as an integer or -1 on failure
func (h *Hash) Version() int {
	if h.version == "" {
		return -1
	}
	if h.version[0] != 'v' {
		return -1
	}
	if spos := strings.IndexRune(h.version[1:], '_'); spos >= 0 {
		n, e := strconv.Atoi(h.version[1 : spos+1])
		if e != nil {
			return -1
		}
		return n
	}
	return -1
}

// HashMD5 takes a data structure and returns a hash string of that data structure at the version asked.
// This function uses md5 hashing function and default formatter. See also Dump() function.
func (h *Hash) HashMD5(c interface{}) string {
	return fmt.Sprintf("v%d_%x", h.Version(), h.Md5(c))
}

// HashSHA256 returns the sha 256 hash of the configurationServiceI in a standard base64 encoded string
func (h *Hash) HashSHA256(c interface{}) []byte {
	sum := sha256.Sum256(h.Dump(c))
	return sum[:]
}

// Dump takes a data structure and returns its byte representation.
// This can be useful if you need to use your own hashing function or formatter.
func (h *Hash) Dump(c interface{}) []byte {
	return serialize(c, h.Version())
}

// Md5 takes a data structure and returns its md5 hash.
// This is a shorthand for md5.Sum(Dump(c, version)).
func (h *Hash) Md5(c interface{}) []byte {
	sum := md5.Sum(h.Dump(c))
	return sum[:]
}

// Sha1 takes a data structure and returns its sha1 hash.
// This is a shorthand for sha1.Sum(Dump(c, version)).
func (h *Hash) Sha1(c interface{}) []byte {
	sum := sha1.Sum(h.Dump(c))
	return sum[:]
}

func (h *Hash) Hash64(v interface{}) (uint64, *errorAVA.Error) {
	// Create our walker and walk the structure
	w := &walker{
		h:       h.opts.Hasher,
		tag:     h.opts.TagName,
		zeronil: h.opts.ZeroNil,
	}
	return w.visit(reflect.ValueOf(v), nil)
}

type item struct {
	name  string
	value reflect.Value
}

type itemSorter []item

func (s itemSorter) Len() int {
	return len(s)
}

func (s itemSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s itemSorter) Less(i, j int) bool {
	return s[i].name < s[j].name
}

type tagError string

func (e tagError) Error() string {
	return "incorrect tag " + string(e)
}

type structFieldFilter func(reflect.StructField, *item) (bool, error)

func writeValue(buf *bytes.Buffer, val reflect.Value, fltr structFieldFilter) {
	switch val.Kind() {
	case reflect.String:
		buf.WriteByte('"')
		buf.WriteString(val.String())
		buf.WriteByte('"')
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		buf.WriteString(strconv.FormatInt(val.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		buf.WriteString(strconv.FormatUint(val.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		buf.WriteString(strconv.FormatFloat(val.Float(), 'E', -1, 64))
	case reflect.Bool:
		if val.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteByte('f')
		}
	case reflect.Ptr:
		if !val.IsNil() || val.Type().Elem().Kind() == reflect.Struct {
			writeValue(buf, reflect.Indirect(val), fltr)
		} else {
			writeValue(buf, reflect.Zero(val.Type().Elem()), fltr)
		}
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		len := val.Len()
		for i := 0; i < len; i++ {
			if i != 0 {
				buf.WriteByte(',')
			}
			writeValue(buf, val.Index(i), fltr)
		}
		buf.WriteByte(']')
	case reflect.Map:
		mk := val.MapKeys()
		items := make([]item, len(mk), len(mk))
		// Get all values
		for i := range items {
			items[i].name = formatValue(mk[i], fltr)
			items[i].value = val.MapIndex(mk[i])
		}

		// Sort values by key
		sort.Sort(itemSorter(items))

		buf.WriteByte('[')
		for i := range items {
			if i != 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(items[i].name)
			buf.WriteByte(':')
			writeValue(buf, items[i].value, fltr)
		}
		buf.WriteByte(']')
	case reflect.Struct:
		vtype := val.Type()
		flen := vtype.NumField()
		items := make([]item, 0, flen)
		// Get all fields
		for i := 0; i < flen; i++ {
			field := vtype.Field(i)
			it := item{field.Name, val.Field(i)}
			if fltr != nil {
				ok, err := fltr(field, &it)
				if err != nil && strings.Contains(err.Error(), "method:") {
					panic(err)
				}
				if !ok {
					continue
				}
			}
			items = append(items, it)
		}
		// Sort fields by name
		sort.Sort(itemSorter(items))

		buf.WriteByte('{')
		for i := range items {
			if i != 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(items[i].name)
			buf.WriteByte(':')
			writeValue(buf, items[i].value, fltr)
		}
		buf.WriteByte('}')
	case reflect.Interface:
		writeValue(buf, reflect.ValueOf(val.Interface()), fltr)
	default:
		buf.WriteString(val.String())
	}
}

func formatValue(val reflect.Value, fltr structFieldFilter) string {
	if val.Kind() == reflect.String {
		return "\"" + val.String() + "\""
	}

	var buf bytes.Buffer
	writeValue(&buf, val, fltr)

	return string(buf.Bytes())
}

func filterField(f reflect.StructField, i *item, version int) (bool, error) {
	var err error
	ver := 0
	lastver := -1
	if str := f.Tag.Get("hash"); str != "" {
		if str == "-" {
			return false, nil
		}
		for _, tag := range strings.Split(str, " ") {
			args := strings.Split(strings.TrimSpace(tag), ":")
			if len(args) != 2 {
				return false, tagError(tag)
			}
			switch args[0] {
			case "name":
				i.name = args[1]
			case "version":
				if ver, err = strconv.Atoi(args[1]); err != nil {
					return false, tagError(tag)
				}
			case "lastversion":
				if lastver, err = strconv.Atoi(args[1]); err != nil {
					return false, tagError(tag)
				}
			case "method":
				property, found := f.Type.MethodByName(strings.TrimSpace(args[1]))
				if !found || property.Type.NumOut() != 1 {
					return false, tagError(tag)
				}
				i.value = property.Func.Call([]reflect.Value{i.value})[0]
			}
		}
	} else {
		if str := f.Tag.Get("lastversion"); str != "" {
			if lastver, err = strconv.Atoi(str); err != nil {
				return false, tagError(str)
			}
		}
		if str := f.Tag.Get("version"); str != "" {
			if ver, err = strconv.Atoi(str); err != nil {
				return false, tagError(str)
			}
		}
	}
	if lastver != -1 && lastver < version {
		return false, nil
	}
	if ver > version {
		return false, nil
	}
	return true, nil
}

func serialize(object interface{}, version int) []byte {
	var buf bytes.Buffer

	writeValue(&buf, reflect.ValueOf(object),
		func(f reflect.StructField, i *item) (bool, error) {
			return filterField(f, i, version)
		})

	return buf.Bytes()
}

type walker struct {
	h       _hash.Hash64
	tag     string
	zeronil bool
}

type visitOpts struct {
	// Flags are a bitmask of flags to affect behavior of this visit
	Flags visitFlag

	// Information about the struct containing this field
	Struct      interface{}
	StructField string
}

// Includable is an interface that can optionally be implemented by a struct. It will be called for each field in the struct to check whether it should be included in the hash.
type Includable interface {
	HashInclude(field string, v interface{}) (bool, *errorAVA.Error)
}

// IncludableMap is an interface that can optionally be implemented by a struct. It will be called when a map-type field is found to ask the struct if the map item should be included in the hash.
type IncludableMap interface {
	HashIncludeMap(field string, k, v interface{}) (bool, *errorAVA.Error)
}

func (w *walker) visit(v reflect.Value, opts *visitOpts) (uint64, *errorAVA.Error) {
	t := reflect.TypeOf(0)

	// Loop since these can be wrapped in multiple layers of pointers and interfaces.
	for {
		// If we have an interface, dereference it. We have to do this up here because it might be a nil in there and the check below must catch that.
		if v.Kind() == reflect.Interface {
			v = v.Elem()
			continue
		}

		if v.Kind() == reflect.Ptr {
			if w.zeronil {
				t = v.Type().Elem()
			}
			v = reflect.Indirect(v)
			continue
		}

		break
	}

	// If it is nil, treat it like a zero.
	if !v.IsValid() {
		v = reflect.Zero(t)
	}

	// Binary writing can use raw ints, we have to convert to a sized-int, we'll choose the largest...
	switch v.Kind() {
	case reflect.Int:
		v = reflect.ValueOf(int64(v.Int()))
	case reflect.Uint:
		v = reflect.ValueOf(uint64(v.Uint()))
	case reflect.Bool:
		var tmp int8
		if v.Bool() {
			tmp = 1
		}
		v = reflect.ValueOf(tmp)
	}

	k := v.Kind()

	// We can shortcut numeric values by directly binary writing them
	if k >= reflect.Int && k <= reflect.Complex64 {
		// A direct hash calculation
		w.h.Reset()
		err := binary.Write(w.h, binary.LittleEndian, v.Interface())
		return w.h.Sum64(), errorHashAVA.WriteBinary(err, v.String())
	}

	switch k {
	case reflect.Array:
		var h uint64
		l := v.Len()
		for i := 0; i < l; i++ {
			current, err := w.visit(v.Index(i), nil)
			if err != nil {
				return 0, err
			}

			h = hashUpdateOrdered(w.h, h, current)
		}

		return h, nil

	case reflect.Map:
		var includeMap IncludableMap
		if opts != nil && opts.Struct != nil {
			if v, ok := opts.Struct.(IncludableMap); ok {
				includeMap = v
			}
		}

		// Build the hash for the map. We do this by XOR-ing all the key and value hashes. This makes it deterministic despite ordering.
		var h uint64
		for _, k := range v.MapKeys() {
			v := v.MapIndex(k)
			if includeMap != nil {
				incl, err := includeMap.HashIncludeMap(
					opts.StructField, k.Interface(), v.Interface())
				if err != nil {
					return 0, err
				}
				if !incl {
					continue
				}
			}

			kh, err := w.visit(k, nil)
			if err != nil {
				return 0, err
			}
			vh, err := w.visit(v, nil)
			if err != nil {
				return 0, err
			}

			fieldHash := hashUpdateOrdered(w.h, kh, vh)
			h = hashUpdateUnordered(h, fieldHash)
		}

		return h, nil

	case reflect.Struct:
		parent := v.Interface()
		var include Includable
		if impl, ok := parent.(Includable); ok {
			include = impl
		}

		t := v.Type()
		h, err := w.visit(reflect.ValueOf(t.Name()), nil)
		if err != nil {
			return 0, err
		}

		l := v.NumField()
		for i := 0; i < l; i++ {
			if innerV := v.Field(i); v.CanSet() || t.Field(i).Name != "_" {
				var f visitFlag
				fieldType := t.Field(i)
				if fieldType.PkgPath != "" {
					// Unexported
					continue
				}

				tag := fieldType.Tag.Get(w.tag)
				if tag == "ignore" || tag == "-" {
					// Ignore this field
					continue
				}

				// if string is set, use the string value
				if tag == "string" {
					if impl, ok := innerV.Interface().(fmt.Stringer); ok {
						innerV = reflect.ValueOf(impl.String())
					} else {
						return 0, errorHashAVA.TypeUnknown(nil, fmt.Sprintf("Field: %s", v.Type().Field(i).Name))
					}
				}

				// Check if we implement includable and check it
				if include != nil {
					incl, err := include.HashInclude(fieldType.Name, innerV)
					if err != nil {
						return 0, err
					}
					if !incl {
						continue
					}
				}

				switch tag {
				case "set":
					f |= visitFlagSet
				}

				kh, err := w.visit(reflect.ValueOf(fieldType.Name), nil)
				if err != nil {
					return 0, err
				}

				vh, err := w.visit(innerV, &visitOpts{
					Flags:       f,
					Struct:      parent,
					StructField: fieldType.Name,
				})
				if err != nil {
					return 0, err
				}

				fieldHash := hashUpdateOrdered(w.h, kh, vh)
				h = hashUpdateUnordered(h, fieldHash)
			}
		}

		return h, nil

	case reflect.Slice:
		// We have two behaviors here. If it isn't a set, then we just
		// visit all the elements. If it is a set, then we do a deterministic
		// hash code.
		var h uint64
		var set bool
		if opts != nil {
			set = (opts.Flags & visitFlagSet) != 0
		}
		l := v.Len()
		for i := 0; i < l; i++ {
			current, err := w.visit(v.Index(i), nil)
			if err != nil {
				return 0, err
			}

			if set {
				h = hashUpdateUnordered(h, current)
			} else {
				h = hashUpdateOrdered(w.h, h, current)
			}
		}

		return h, nil

	case reflect.String:
		// Directly hash
		w.h.Reset()
		if _, err := w.h.Write([]byte(v.String())); err != nil {
			return w.h.Sum64(), errorHashAVA.WriteBinary(err, v.String())
		}
		return w.h.Sum64(), nil
	default:
		return 0, errorHashAVA.TypeUnknown(nil, fmt.Sprintf("Unknown kind to hash: %s", k))
	}

}

func hashUpdateOrdered(h _hash.Hash64, a, b uint64) uint64 {
	// For ordered updates, use a real hash function
	h.Reset()

	// We just panic if the binary writes fail because we are writing
	// an int64 which should never be fail-able.
	e1 := binary.Write(h, binary.LittleEndian, a)
	e2 := binary.Write(h, binary.LittleEndian, b)
	if e1 != nil {
		panic(e1)
	}
	if e2 != nil {
		panic(e2)
	}

	return h.Sum64()
}

func hashUpdateUnordered(a, b uint64) uint64 {
	return a ^ b
}

// visitFlag is used as a bitmask for affecting visit behavior
type visitFlag uint

const (
	visitFlagInvalid visitFlag = iota
	visitFlagSet               = iota << 1
)
