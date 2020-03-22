package error_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type subgroupUnknownSuite struct {
	suite.Suite
}

func TestSubgroupUnknownInit(t *testing.T) {
	suite.Run(t, new(subgroupUnknownSuite))
}

func (r *subgroupUnknownSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *subgroupUnknownSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *subgroupUnknownSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *subgroupUnknownSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *subgroupUnknownSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *subgroupUnknownSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *subgroupUnknownSuite) TestSubgroupTypeUnknown() {
	Convey("Given a AVA SubgroupTypeUnknown error", r.T(), func() {
		Convey("Went all fields are OK ", func() {
			details := "Details"
			err := errorAVA.SubgroupTypeUnknown(nil, details)

			So(err, ShouldNotBeNil)

			So(err.Group, ShouldEqual, errorAVA.GroupUnknown)
			So(err.Subgroup, ShouldEqual, errorAVA.SubgroupUnknown)
			So(err.Code, ShouldEqual, errorAVA.StatusSubgroupUnknown)

			So(err.Err, ShouldBeNil)

			So(err.Message, ShouldResemble, errorAVA.StatusTextFunc(errorAVA.StatusSubgroupUnknown))
			So(err.Details, ShouldResemble, fmt.Sprintf("%v.", details))

			So(err.Info, ShouldResemble, &errorAVA.CallInfo{
				PackageName: "github.com/ver13/ava/test/common/error_test",
				FileName:    "subGroupTypeUnknown_test.go",
				FuncName:    "TestSubgroupTypeUnknown",
				Line:        49,
			})
		})
		Convey("Went all functions are OK ", func() {
			details := "Details"
			err := errorAVA.SubgroupTypeUnknown(nil, details)

			So(err, ShouldNotBeNil)

			errorStr := err.String()
			json, errorJSON := err.ToJSON()
			So(errorJSON, ShouldBeNil)
			So(json, ShouldResemble, "{\n    \"Group\": 26,\n    \"Subgroup\": 17,\n    \"Err\": null,\n    \"Code\": 2,\n    \"Message\": \"Subgroup type unknown.\",\n    \"Details\": \"Details.\",\n    \"Info\": {\n        \"PackageName\": \"github.com/ver13/ava/test/common/error_test\",\n        \"FileName\": \"subGroupTypeUnknown_test.go\",\n        \"FuncName\": \"TestSubgroupTypeUnknown\",\n        \"Line\": 71\n    }\n}")
			So(errorStr, ShouldResemble, "Group: Unknown, Subgroup: Subgroup(17), Details: Details., Error: <nil>, Info: &{github.com/ver13/ava/test/common/error_test subGroupTypeUnknown_test.go TestSubgroupTypeUnknown 71}")
		})
		Convey("Went it's wrong ", func() {
			details := "Details"
			err := errorAVA.SubgroupTypeUnknown(nil, details)

			So(err, ShouldNotBeNil)
			So(err.Details, ShouldNotResemble, details)
		})
	})
}
