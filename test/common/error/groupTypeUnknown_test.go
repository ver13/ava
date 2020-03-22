package error_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type groupUnknownSuite struct {
	suite.Suite
}

func TestGroupUnknownInit(t *testing.T) {
	suite.Run(t, new(groupUnknownSuite))
}

func (r *groupUnknownSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *groupUnknownSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *groupUnknownSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *groupUnknownSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *groupUnknownSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *groupUnknownSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *groupUnknownSuite) TestGroupTypeUnknown() {
	Convey("Given a AVA GroupTypeUnknown error", r.T(), func() {
		Convey("Went all fields are OK ", func() {
			details := "Details"
			err := errorAVA.GroupTypeUnknown(nil, details)

			So(err, ShouldNotBeNil)

			So(err.Group, ShouldEqual, errorAVA.GroupUnknown)
			So(err.Subgroup, ShouldEqual, errorAVA.SubgroupUnknown)
			So(err.Code, ShouldEqual, errorAVA.GroupUnknownCode)

			So(err.Err, ShouldBeNil)

			So(err.Message, ShouldResemble, errorAVA.StatusTextFunc(errorAVA.GroupUnknownCode))
			So(err.Details, ShouldResemble, fmt.Sprintf("%v.", details))

			So(err.Info, ShouldResemble, &errorAVA.CallInfo{
				PackageName: "github.com/ver13/ava/test/common/error_test",
				FileName:    "groupTypeUnknown_test.go",
				FuncName:    "TestGroupTypeUnknown",
				Line:        49,
			})
		})
		Convey("Went all functions are OK ", func() {
			details := "Details"
			err := errorAVA.GroupTypeUnknown(nil, details)

			So(err, ShouldNotBeNil)

			errorStr := err.String()
			json, errorJSON := err.ToJSON()

			So(errorJSON, ShouldBeNil)
			So(json, ShouldResemble, "{\n    \"Group\": 26,\n    \"Subgroup\": 17,\n    \"Err\": null,\n    \"Code\": 1,\n    \"Message\": \"Group type unknown.\",\n    \"Details\": \"Details.\",\n    \"Info\": {\n        \"PackageName\": \"github.com/ver13/ava/test/common/error_test\",\n        \"FileName\": \"groupTypeUnknown_test.go\",\n        \"FuncName\": \"TestGroupTypeUnknown\",\n        \"Line\": 71\n    }\n}")
			So(errorStr, ShouldResemble, "Group: Unknown, Subgroup: Subgroup(17), Details: Details., Error: <nil>, Info: &{github.com/ver13/ava/test/common/error_test groupTypeUnknown_test.go TestGroupTypeUnknown 71}")
		})
		Convey("Went it's wrong ", func() {
			details := "Details"
			err := errorAVA.GroupTypeUnknown(nil, details)

			So(err, ShouldNotBeNil)
			So(err.Details, ShouldNotResemble, details)
		})
	})
}
