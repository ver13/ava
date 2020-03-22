package general_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	generalErrorAVA "github.com/ver13/ava/pkg/common/error/general"
)

type isNilSuite struct {
	suite.Suite
}

func TestIsNilInit(t *testing.T) {
	suite.Run(t, new(isNilSuite))
}

func (r *isNilSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *isNilSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *isNilSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *isNilSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *isNilSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *isNilSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *isNilSuite) TestIsNil() {
	Convey("Given a AVA error", r.T(), func() {
		Convey("Went all fields are OK ", func() {
			details := "Details"
			err := generalErrorAVA.IsNil(nil, details)

			So(err, ShouldNotBeNil)

			So(err.Group, ShouldEqual, errorAVA.GroupGeneral)
			So(err.Subgroup, ShouldEqual, errorAVA.SubgroupGeneral)
			So(err.Code, ShouldEqual, generalErrorAVA.IsNilCode)

			So(err.Err, ShouldBeNil)

			So(err.Message, ShouldResemble, generalErrorAVA.StatusTextFunc(generalErrorAVA.IsNilCode))
			So(err.Details, ShouldResemble, fmt.Sprintf("%v.", details))

			So(err.Info, ShouldResemble, &errorAVA.CallInfo{
				PackageName: "github.com/ver13/ava/test/common/error/general_test",
				FileName:    "isNil_test.go",
				FuncName:    "TestIsNil",
				Line:        50,
			})
		})
		Convey("Went all functions are OK ", func() {
			details := "Details"
			err := generalErrorAVA.IsNil(nil, details)

			So(err, ShouldNotBeNil)

			errorStr := err.String()
			json, errorJSON := err.ToJSON()

			So(errorJSON, ShouldBeNil)
			So(json, ShouldResemble, "{\n    \"Group\": 0,\n    \"Subgroup\": 0,\n    \"Err\": null,\n    \"Code\": 2,\n    \"Message\": \"ErrorHTTP is NIL.\",\n    \"Details\": \"Details.\",\n    \"Info\": {\n        \"PackageName\": \"github.com/ver13/ava/test/common/error/general_test\",\n        \"FileName\": \"isNil_test.go\",\n        \"FuncName\": \"TestIsNil\",\n        \"Line\": 72\n    }\n}")
			So(errorStr, ShouldResemble, "Group: General, Subgroup: General, Details: Details., Error: <nil>, Info: &{github.com/ver13/ava/test/common/error/general_test isNil_test.go TestIsNil 72}")
		})
		Convey("Went it's wrong ", func() {
			details := "Details"
			err := generalErrorAVA.IsNil(nil, details)

			So(err, ShouldNotBeNil)
			So(err.Details, ShouldNotResemble, details)
		})
	})
}
