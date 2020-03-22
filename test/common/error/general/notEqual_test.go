package general_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	generalErrorAVA "github.com/ver13/ava/pkg/common/error/general"
)

type notEqualSuite struct {
	suite.Suite
}

func TestNotEqualInit(t *testing.T) {
	suite.Run(t, new(notEqualSuite))
}

func (r *notEqualSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *notEqualSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *notEqualSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *notEqualSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *notEqualSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *notEqualSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *notEqualSuite) TestNotEqual() {
	Convey("Given a AVA error", r.T(), func() {
		Convey("Went all fields are OK ", func() {
			details := "Details"
			err := generalErrorAVA.NotEqual(nil, details)

			So(err, ShouldNotBeNil)

			So(err.Group, ShouldEqual, errorAVA.GroupGeneral)
			So(err.Subgroup, ShouldEqual, errorAVA.SubgroupGeneral)
			So(err.Code, ShouldEqual, generalErrorAVA.NotEqualCode)

			So(err.Err, ShouldBeNil)

			So(err.Message, ShouldResemble, generalErrorAVA.StatusTextFunc(generalErrorAVA.NotEqualCode))
			So(err.Details, ShouldResemble, fmt.Sprintf("%v.", details))

			So(err.Info, ShouldResemble, &errorAVA.CallInfo{
				PackageName: "github.com/ver13/ava/test/common/error/general_test",
				FileName:    "notEqual_test.go",
				FuncName:    "TestNotEqual",
				Line:        50,
			})
		})
		Convey("Went all functions are OK ", func() {
			details := "Details"
			err := generalErrorAVA.NotEqual(nil, details)

			So(err, ShouldNotBeNil)

			errorStr := err.String()
			json, errorJSON := err.ToJSON()
			So(errorJSON, ShouldBeNil)
			So(json, ShouldResemble, "{\n    \"Group\": 0,\n    \"Subgroup\": 0,\n    \"Err\": null,\n    \"Code\": 1,\n    \"Message\": \"Error not equal.\",\n    \"Details\": \"Details.\",\n    \"Info\": {\n        \"PackageName\": \"github.com/ver13/ava/test/common/error/general_test\",\n        \"FileName\": \"notEqual_test.go\",\n        \"FuncName\": \"TestNotEqual\",\n        \"Line\": 72\n    }\n}")
			So(errorStr, ShouldResemble, "Group: General, Subgroup: General, Details: Details., Error: <nil>, Info: &{github.com/ver13/ava/test/common/error/general_test notEqual_test.go TestNotEqual 72}")
		})
		Convey("Went it's wrong ", func() {
			details := "Details"
			err := generalErrorAVA.NotEqual(nil, details)

			So(err, ShouldNotBeNil)
			So(err.Details, ShouldNotResemble, details)
		})
	})
}
