package hash_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/hash"
)

type hashSuite struct {
	suite.Suite

	hashv1 HashI
	hashv2 HashI
	hashv3 HashI
}

func TestHashInit(t *testing.T) {
	suite.Run(t, new(hashSuite))
}

func (r *hashSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *hashSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *hashSuite) SetupSuite() {
	r.T().Log("SetupSuite")
	r.hashv1 = NewHash("v1")
	r.hashv2 = NewHash("v2")
	r.hashv3 = NewHash("v3")
}

func (r *hashSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *hashSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *hashSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

type First struct {
	Bool   bool    `version:"1"`
	String string  `version:"2"`
	Int    int     `version:"1" lastversion:"1"`
	Float  float64 `version:"1"`
	Struct *Second `version:"1"`
	Uint   uint    `version:"1"`
}

type Second struct {
	Map   map[string]string `version:"1"`
	Slice []int             `version:"1"`
}

type Tags1 struct {
	Int int    `hash:"-"`
	Str string `hash:"name:Foo version:1 lastversion:2"`
	Bar string `hash:"version:1"`
}

type Tags2 struct {
	Foo string
	Bar string
}

type Tags3 struct {
	Bar string
}

type Tags4 struct {
	Data1 ambiguousData `hash:"method:Serialize"`
	Data2 ambiguousData `hash:"method:Normalize"`
}

type Tags5 struct {
	Data1 ambiguousData `hash:"method:UnknownMethod"`
}

type Nils struct {
	Str   *string
	Int   *int
	Bool  *bool
	Map   map[string]string
	Slice []string
}

type unexportedTags struct {
	foo  string
	bar  string
	aMap map[string]string
}

type interfaceStruct struct {
	Name            string
	Interface1      interface{}
	InterfaceIgnore interface{} `hash:"-"`
}

type ambiguousData struct {
	Prefix string
	Suffix string
}

func (p ambiguousData) Serialize() string {
	return p.Prefix + p.Suffix
}

func (p ambiguousData) Normalize() ambiguousData {
	return ambiguousData{p.Prefix + p.Suffix, ""}
}

func dataSetup() *First {
	tmpmap := make(map[string]string)
	tmpmap["foo"] = "bar"
	tmpmap["baz"] = "go"
	tmpslice := make([]int, 3)
	tmpslice[0] = 0
	tmpslice[1] = 1
	tmpslice[2] = 2
	return &First{
		Bool:   true,
		String: "test",
		Int:    123456789,
		Float:  65.3458,
		Struct: &Second{
			Map:   tmpmap,
			Slice: tmpslice,
		},
		Uint: 1,
	}
}

func (r *hashSuite) TestHash() {
	Convey("Given hash ", r.T(), func() {
		v1Hash := "v1_e8e67581aee36d7237603381a9cbd9fc"
		v2Hash := "v2_5e51490d7c24c4b7a9e63c04f55734eb"

		data := dataSetup()
		v1 := r.hashv1.HashMD5(data)
		So(v1, ShouldNotResemble, v1Hash)

		v2 := r.hashv2.HashMD5(data)
		So(v2, ShouldNotResemble, v2Hash)

		v1md5 := fmt.Sprintf("v1_%x", r.hashv1.Md5(data))
		So(v1md5, ShouldNotResemble, v1Hash)

		v2md5 := fmt.Sprintf("v2_%x", r.hashv2.Md5(data))
		So(v2md5, ShouldNotResemble, v2Hash)
	})
}

func (r *hashSuite) TestHash_Tags() {
	Convey("Given hash ", r.T(), func() {
		t1 := Tags1{11, "foo", "bar"}
		t1x := Tags1{22, "foo", "bar"}
		t2 := Tags2{"foo", "bar"}
		t3 := Tags3{"bar"}

		t1Dump := string(r.hashv1.Dump(t1))
		t1xDump := string(r.hashv1.Dump(t1x))
		So(t1Dump, ShouldResemble, t1xDump)

		t2Dump := string(r.hashv1.Dump(t2))
		So(t1Dump, ShouldResemble, t2Dump)

		t1v3Dump := string(r.hashv3.Dump(t1))
		t3v3Dump := string(r.hashv3.Dump(t3))
		So(t1v3Dump, ShouldResemble, t3v3Dump)
	})
}

func (r *hashSuite) TestHash_Nils() {
	Convey("Given hash ", r.T(), func() {
		s1 := Nils{
			Str:   nil,
			Int:   nil,
			Bool:  nil,
			Map:   nil,
			Slice: nil,
		}

		s2 := Nils{
			Str:   new(string),
			Int:   new(int),
			Bool:  new(bool),
			Map:   make(map[string]string),
			Slice: make([]string, 0),
		}

		s1Dump := string(r.hashv1.Dump(s1))
		s2Dump := string(r.hashv1.Dump(s2))
		So(s1Dump, ShouldResemble, s2Dump)
	})
}

func (r *hashSuite) TestHash_UnexportedFields() {
	Convey("Given hash ", r.T(), func() {
		v1Hash := "v1_750efb7c919caf87f2ab0d119650c87d"
		data := unexportedTags{
			foo: "foo",
			bar: "bar",
			aMap: map[string]string{
				"key1": "val",
			},
		}
		v1 := r.hashv1.HashMD5(data)
		So(v1, ShouldNotResemble, v1Hash)

		v1md5 := fmt.Sprintf("v1_%x", NewHash("v1").Md5(data))
		So(v1md5, ShouldNotResemble, v1Hash)
	})
}

func (r *hashSuite) TestHash_InterfaceField() {
	Convey("Given hash ", r.T(), func() {
		a := interfaceStruct{
			Name:            "name",
			Interface1:      "a",
			InterfaceIgnore: "b",
		}

		b := interfaceStruct{
			Name:            "name",
			Interface1:      "b",
			InterfaceIgnore: "b",
		}

		c := interfaceStruct{
			Name:            "name",
			Interface1:      "b",
			InterfaceIgnore: "c",
		}

		ha := r.hashv1.HashMD5(a)
		hb := r.hashv1.HashMD5(b)
		hc := r.hashv1.HashMD5(c)

		So(ha, ShouldResemble, hb)
		So(hb, ShouldResemble, hc)

		b.Interface1 = map[string]string{"key": "value"}
		c.Interface1 = map[string]string{"key": "value"}

		hb = r.hashv1.HashMD5(b)
		hc = r.hashv1.HashMD5(c)

		So(hb, ShouldResemble, hc)

		c.Interface1 = map[string]string{"key1": "value1"}
		hc = r.hashv1.HashMD5(c)

		So(hb, ShouldResemble, hc)
	})
}

func (r *hashSuite) TestHash_Method() {
	Convey("Given hash ", r.T(), func() {
		dump1 := r.hashv1.Dump(Tags4{
			Data1: ambiguousData{"123", "45"},
			Data2: ambiguousData{"12", "345"},
		})
		dump2 := r.hashv1.Dump(Tags4{
			Data1: ambiguousData{"12", "345"},
			Data2: ambiguousData{"123", "45"},
		})

		So(dump1, ShouldResemble, dump2)
	})
}

func (r *hashSuite) TestHash_MethodPanic() {
	Convey("Given hash ", r.T(), func() {
		defer func() {
			r := recover()
			So(r, ShouldNotBeNil)
		}()
		_ = r.hashv1.Dump(Tags5{Data1: ambiguousData{"123", "45"}})
	})
}
