package string

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorStringAVA "github.com/ver13/ava/pkg/common/string/error"
)

type String string

var (
	s String

	once sync.Once
)

func init() {
	once.Do(func() {
		// No hay que crear el objeto, porque es de un tipo simple
	})
}

func GetInstance() *String {
	return &s
}

// MarshalJSON marshals data to an indented string.
func (s *String) MarshalJSON(data interface{}, indent string) string {
	buffer, err := json.MarshalIndent(data, "", indent)
	if err != nil {
		return ""
	}
	return string(buffer)
}
func StringMarshalJSON(data interface{}, indent string) string {
	return GetInstance().MarshalJSON(data, indent)
}

func (s *String) ListContains(l []string, str string) bool {
	for i := range l {
		if l[i] == str {
			return true
		}
	}
	return false
}
func StringListContains(l []string, s string) bool {
	return GetInstance().ListContains(l, s)
}

func (s *String) ListContainsCaseInsensitive(l []string, str string) bool {
	tmp := strings.ToLower(str)
	s = (*String)(&tmp)
	for i := range l {
		if strings.ToLower(l[i]) == str {
			return true
		}
	}
	return false
}
func StringListContainsCaseInsensitive(l []string, s string) bool {
	return GetInstance().ListContainsCaseInsensitive(l, s)
}

func (s *String) PrettifyJSON(compactJSON string) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(compactJSON), "", "\t")
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
func StringPrettifyJSON(compactJSON string) string {
	return GetInstance().PrettifyJSON(compactJSON)
}

func (s *String) EscapeJSON(jsonString string) string {
	jsonString = strings.Replace(jsonString, `\`, `\\`, -1)
	jsonString = strings.Replace(jsonString, `"`, `\"`, -1)
	return jsonString
}
func StringEscapeJSON(jsonString string) string {
	return GetInstance().EscapeJSON(jsonString)
}

// StripHTMLTags strips HTML/XML tags from text.
func (s *String) StripHTMLTags(text string) (plainText string) {
	var buf *bytes.Buffer
	tagClose := -1
	tagStart := -1
	for i, char := range text {
		if char == '<' {
			if buf == nil {
				buf = bytes.NewBufferString(text)
				buf.Reset()
			}
			buf.WriteString(text[tagClose+1 : i])
			tagStart = i
		} else if char == '>' && tagStart != -1 {
			tagClose = i
			tagStart = -1
		}
	}
	if buf == nil {
		return text
	}
	buf.WriteString(text[tagClose+1:])
	return buf.String()
}
func StringStripHTMLTags(text string) (plainText string) {
	return GetInstance().StripHTMLTags(text)
}

// ReplaceHTMLTags replaces HTML/XML tags from text with replacement.
func (s *String) ReplaceHTMLTags(text, replacement string) (plainText string) {
	var buf *bytes.Buffer
	tagClose := -1
	tagStart := -1
	for i, char := range text {
		if char == '<' {
			if buf == nil {
				buf = bytes.NewBufferString(text)
				buf.Reset()
			}
			buf.WriteString(text[tagClose+1 : i])
			tagStart = i
		} else if char == '>' && tagStart != -1 {
			buf.WriteString(replacement)
			tagClose = i
			tagStart = -1
		}
	}
	if buf == nil {
		return text
	}
	buf.WriteString(text[tagClose+1:])
	return buf.String()
}
func StringReplaceHTMLTags(text, replacement string) (plainText string) {
	return GetInstance().StripHTMLTags(text)
}

// MD5Hex returns the hex encoded MD5 hash of data
func (s *String) MD5Hex(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func StringMD5Hex(data string) string {
	return GetInstance().MD5Hex(data)
}

// SHA1Base64 returns the base64 encoded SHA1 hash of data
func (s *String) SHA1Base64(data string) string {
	hash := sha1.Sum([]byte(data))
	return base64.StdEncoding.EncodeToString(hash[:])
}
func StringSHA1Base64(data string) string {
	return GetInstance().SHA1Base64(data)
}

func (s *String) AddURLParam(url, name, value string) string {
	var separator string
	if strings.IndexRune(url, '?') == -1 {
		separator = "?"
	} else {
		separator = "&"
	}
	return url + separator + name + "=" + value
}
func StringAddURLParam(url, name, value string) string {
	return GetInstance().AddURLParam(url, name, value)
}

func (s *String) ConvertTime(timeString string, formatIn string, formatOut string) (string, *errorAVA.Error) {
	if timeString == "" {
		return "", nil
	}
	t, err := time.Parse(formatIn, timeString)
	if err != nil {
		return "", errorStringAVA.TimeParser(err, fmt.Sprintf("format: [%s] - TimeGmf: [%s]", formatIn, timeString))
	}
	return t.Format(formatOut), nil
}
func StringConvertTime(timeString string, formatIn string, formatOut string) (string, *errorAVA.Error) {
	return GetInstance().ConvertTime(timeString, formatIn, formatOut)
}

func (s *String) CSV(records [][]string) string {
	var b strings.Builder
	writer := csv.NewWriter(&b)
	err := writer.WriteAll(records)
	if err != nil {
		return ""
	}
	return b.String()
}
func StringCSV(records [][]string) string {
	return GetInstance().CSV(records)
}

func (s *String) ToInt(str string) int {
	i, _ := strconv.ParseInt(str, 10, 64)
	return int(i)
}
func StringToInt(s string) int {
	return GetInstance().ToInt(s)
}

func (s *String) ToFloat(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
func StringToFloat(s string) float64 {
	return GetInstance().ToFloat(s)
}

func (s *String) ToBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}
func StringToBool(s string) bool {
	return GetInstance().ToBool(s)
}

func (s *String) InSlice(str string, slice []string) bool {
	for i := range slice {
		if slice[i] == str {
			return true
		}
	}
	return false
}
func StringInSlice(s string, slice []string) bool {
	return GetInstance().InSlice(s, slice)
}

// JoinFormat formats every value in values with format and joins the result with sep as separator. values must be a slice of a formatable type
func (s *String) JoinFormat(format string, values interface{}, sep string) string {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Slice {
		panic("values is not a slice")
	}
	var buffer bytes.Buffer
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(fmt.Sprintf(format, v.Index(i).Interface()))
	}
	return buffer.String()
}
func StringJoinFormat(format string, values interface{}, sep string) string {
	return GetInstance().JoinFormat(format, values, sep)
}

// Join formats every value in values according to its default formatting and joins the result with sep as separator. values must be a slice of a formatable type
func (s *String) Join(values interface{}, sep string) string {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Slice {
		panic("values is not a slice")
	}
	var buffer bytes.Buffer
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(fmt.Sprint(v.Index(i).Interface()))
	}
	return buffer.String()
}
func StringJoin(values interface{}, sep string) string {
	return GetInstance().Join(values, sep)
}

func (s *String) FormatBigInt(mem uint64) string {
	switch {
	case mem >= 10e12:
		return fmt.Sprintf("%dT", mem/1e12)
	case mem >= 1e12:
		return strings.TrimSuffix(fmt.Sprintf("%.1fT", float64(mem)/1e12), ".0")

	case mem >= 10e9:
		return fmt.Sprintf("%dG", mem/1e9)
	case mem >= 1e9:
		return strings.TrimSuffix(fmt.Sprintf("%.1fG", float64(mem)/1e9), ".0")

	case mem >= 10e6:
		return fmt.Sprintf("%dM", mem/1e6)
	case mem >= 1e6:
		return strings.TrimSuffix(fmt.Sprintf("%.1fM", float64(mem)/1e6), ".0")

	case mem >= 10e3:
		return fmt.Sprintf("%dk", mem/1e3)
	case mem >= 1e3:
		return strings.TrimSuffix(fmt.Sprintf("%.1fk", float64(mem)/1e3), ".0")
	}
	return fmt.Sprintf("%d", mem)
}
func StringFormatBigInt(mem uint64) string {
	return GetInstance().FormatBigInt(mem)
}

func (s *String) FormatMemory(mem uint64) string {
	return s.FormatBigInt(mem) + "B"
}
func StringFormatMemory(mem uint64) string {
	return GetInstance().FormatMemory(mem)
}

func (s *String) ReplaceMulti(str string, fromTo ...string) string {
	if len(fromTo)%2 != 0 {
		panic("Need even number of fromTo arguments")
	}
	for i := 0; i < len(fromTo); i += 2 {
		str = strings.Replace(str, fromTo[i], fromTo[i+1], -1)
	}
	return str
}
func StringReplaceMulti(s string, fromTo ...string) string {
	return GetInstance().ReplaceMulti(s, fromTo...)
}

func (s *String) ToUpperCamelCase(str string) string {
	var b strings.Builder
	var last byte = '_'
	for _, c := range []byte(str) {
		if c != '_' {
			if last == '_' {
				c = byte(unicode.ToUpper(rune(c)))
			} else {
				c = byte(unicode.ToLower(rune(c)))
			}
			b.WriteByte(c)
		}
		last = c
	}
	return b.String()
}
func StringToUpperCamelCase(str string) string {
	return GetInstance().ToUpperCamelCase(str)
}

func (s *String) ToLowerCamelCase(str string) string {
	var b strings.Builder
	var last byte
	for _, c := range []byte(str) {
		if c != '_' {
			if last == '_' {
				c = byte(unicode.ToUpper(rune(c)))
			} else {
				c = byte(unicode.ToLower(rune(c)))
			}
			b.WriteByte(c)
		}
		last = c
	}
	return b.String()
}
func StringToLowerCamelCase(str string) string {
	return GetInstance().ToLowerCamelCase(str)
}

func (s *String) ToLower(str string) string {
	return strings.ToLower(str)
}
func StringToLower(str string) string {
	return GetInstance().ToLower(str)
}

func (s *String) ToUpper(str string) string {
	return strings.ToUpper(str)
}
func StringToUpper(str string) string {
	return GetInstance().ToUpper(str)
}

func (s *String) MapSortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func StringMapSortedKeys(m map[string]string) []string {
	return GetInstance().MapSortedKeys(m)
}

func (s *String) MapGroupedNumberPostfixSortedKeys(m map[string]string) []string {
	keys := make(StringGroupedNumberPostfixSorter, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Sort(keys)
	return keys
}
func StringMapGroupedNumberPostfixSortedKeys(m map[string]string) []string {
	return GetInstance().MapGroupedNumberPostfixSortedKeys(m)
}

func (s *String) MapGroupedNumberPostfixSortedValues(m map[string]string) []string {
	values := make(StringGroupedNumberPostfixSorter, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	sort.Sort(values)
	return values
}
func StringMapGroupedNumberPostfixSortedValues(m map[string]string) []string {
	return GetInstance().MapGroupedNumberPostfixSortedValues(m)
}

func (s *String) EndsWithNumber(str string) bool {
	if str == "" {
		return false
	}
	c := str[len(str)-1]
	return c >= '0' && c <= '9'
}
func StringEndsWithNumber(str string) bool {
	return GetInstance().EndsWithNumber(str)
}

func (s *String) SplitNumberPostfix(str string) (base, number string) {
	if str == "" {
		return "", ""
	}
	for i := len(str) - 1; i >= 0; i-- {
		c := str[i]
		if c < '0' || c > '9' {
			if i == len(str)-1 {
				return str, ""
			}
			return str[:i+1], str[i+1:]
		}
	}
	return "", str
}
func StringSplitNumberPostfix(str string) (base, number string) {
	return GetInstance().SplitNumberPostfix(str)
}

func (s *String) SplitOnce(str, sep string) (pre, post string) {
	parts := strings.SplitN(str, sep, 1)
	if len(parts) == 2 {
		return parts[0], parts[1]
	} else {
		return parts[0], ""
	}
}
func StringSplitOnce(str, sep string) (pre, post string) {
	return GetInstance().SplitOnce(str, sep)
}

func (s *String) SplitOnceChar(str string, sep byte) (pre, post string) {
	i := strings.IndexByte(str, sep)
	if i == -1 {
		return str, ""
	}
	return str[:i], str[i+1:]
}
func StringSplitOnceChar(str string, sep byte) (pre, post string) {
	return GetInstance().SplitOnceChar(str, sep)
}

func (s *String) SplitOnceRune(str string, sep rune) (pre, post string) {
	sepIndex := -1
	postSepIndex := -1
	for i, c := range str {
		if sepIndex != -1 {
			postSepIndex = i
			break // we got the index after the sep rune
		}
		if c == sep {
			sepIndex = i
			// continue to get index after the current UTF8 rune
		}
	}
	if sepIndex == -1 {
		return str, ""
	}
	return str[:sepIndex], str[postSepIndex:]
}
func StringSplitOnceRune(str string, sep rune) (pre, post string) {
	return GetInstance().SplitOnceRune(str, sep)
}

// Map a function on each element of a slice of strings.
func (s *String) MapFunc(f func(string) string, data []string) []string {
	size := len(data)
	result := make([]string, size, size)
	for i := 0; i < size; i++ {
		result[i] = f(data[i])
	}
	return result
}
func StringMapFunc(f func(string) string, data []string) []string {
	return GetInstance().MapFunc(f, data)
}

// Filter out all strings where the function does not return true.
func (s *String) Filter(f func(string) bool, data []string) []string {
	result := make([]string, 0, 0)
	for _, element := range data {
		if f(element) {
			result = append(result, element)
		}
	}
	return result
}
func StringFilter(f func(string) bool, data []string) []string {
	return GetInstance().Filter(f, data)
}

// FindBetween returns the string between the first occurrences of the tokens start and stop.
// The remainder of the string after the stop token will be returned if found.
// If the tokens couldn't be found, then the whole string will be returned as remainder.
func (s *String) FindBetween(str, start, stop string) (between, remainder string, found bool) {
	begin := strings.Index(str, start)
	if begin == -1 {
		return "", str, false
	}
	between = str[begin+len(start):]
	end := strings.Index(between, stop)
	if end == -1 {
		return "", str, false
	}
	return between[:end], str[begin+len(start)+end+len(stop):], true
}
func StringFindBetween(str, start, stop string) (between, remainder string, found bool) {
	return GetInstance().FindBetween(str, start, stop)
}

// Find returns in found if token has been found in s,
// and returns the remaining string afte token in remainder.
// The whole string s will be returned if found is false.
func (s *String) Find(str, token string) (remainder string, found bool) {
	i := strings.Index(str, token)
	if i == -1 {
		return str, false
	}
	return str[i+len(token):], true
}
func StringFind(str, token string) (remainder string, found bool) {
	return GetInstance().Find(str, token)
}
