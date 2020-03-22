package general_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	generalErrorAVA "github.com/ver13/ava/pkg/common/error/general"
)

type unknownErroruite struct {
	suite.Suite
}

func TestUnknownErrorInit(t *testing.T) {
	suite.Run(t, new(unknownErroruite))
}

func (r *unknownErroruite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *unknownErroruite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *unknownErroruite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *unknownErroruite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *unknownErroruite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *unknownErroruite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *unknownErroruite) TestUnknownError() {
	Convey("Given a AVA error", r.T(), func() {
		Convey("Went all fields are OK ", func() {
			details := "Details"
			err := generalErrorAVA.UnknownError(nil, details)

			So(err, ShouldNotBeNil)

			So(err.Group, ShouldEqual, errorAVA.GroupGeneral)
			So(err.Subgroup, ShouldEqual, errorAVA.SubgroupGeneral)
			So(err.Code, ShouldEqual, generalErrorAVA.UnknownErrorCode)

			So(err.Err, ShouldBeNil)

			So(err.Message, ShouldResemble, generalErrorAVA.StatusTextFunc(generalErrorAVA.UnknownErrorCode))
			So(err.Details, ShouldResemble, fmt.Sprintf("%v.", details))

			So(err.Info, ShouldResemble, &errorAVA.CallInfo{
				PackageName: "github.com/ver13/ava/test/common/error/general_test",
				FileName:    "unknownError_test.go",
				FuncName:    "TestUnknownError",
				Line:        50,
			})
		})
		Convey("Went all functions are OK ", func() {
			details := "Details"
			err := generalErrorAVA.UnknownError(nil, details)

			So(err, ShouldNotBeNil)

			str := err.String()
			json, errorJSON := err.ToJSON()

			So(str, ShouldResemble, "Group: General, Subgroup: General, Details: Details., Error: <nil>, Info: &{github.com/ver13/ava/test/common/error/general_test unknownError_test.go TestUnknownError 72}")
			So(errorJSON, ShouldBeNil)
			So(json, ShouldResemble, "{\n    \"Group\": 0,\n    \"Subgroup\": 0,\n    \"Err\": null,\n    \"Code\": 4,\n    \"Message\": \"Unknown error.\",\n    \"Details\": \"Details.\",\n    \"Info\": {\n        \"PackageName\": \"github.com/ver13/ava/test/common/error/general_test\",\n        \"FileName\": \"unknownError_test.go\",\n        \"FuncName\": \"TestUnknownError\",\n        \"Line\": 72\n    }\n}")
		})
		Convey("Went it's wrong ", func() {
			details := "Details"
			err := generalErrorAVA.UnknownError(nil, details)

			So(err, ShouldNotBeNil)
			So(err.Details, ShouldNotResemble, details)
		})
	})
}
