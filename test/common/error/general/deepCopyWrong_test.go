package general_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	generalErrorAVA "github.com/ver13/ava/pkg/common/error/general"
)

type deepCopyWrongSuite struct {
	suite.Suite
}

func TestDeepCopyWrongInit(t *testing.T) {
	suite.Run(t, new(deepCopyWrongSuite))
}

func (r *deepCopyWrongSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *deepCopyWrongSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *deepCopyWrongSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *deepCopyWrongSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *deepCopyWrongSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *deepCopyWrongSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *deepCopyWrongSuite) TestDeepCopyWrong() {
	Convey("Given a AVA error", r.T(), func() {
		Convey("Went all fields are OK ", func() {
			details := "Details"
			err := generalErrorAVA.DeepCopyWrong(nil, details)

			So(err, ShouldNotBeNil)

			So(err.Group, ShouldEqual, errorAVA.GroupGeneral)
			So(err.Subgroup, ShouldEqual, errorAVA.SubgroupGeneral)
			So(err.Code, ShouldEqual, generalErrorAVA.DeepCopyWrongCode)

			So(err.Err, ShouldBeNil)

			So(err.Message, ShouldResemble, generalErrorAVA.StatusTextFunc(generalErrorAVA.DeepCopyWrongCode))
			So(err.Details, ShouldResemble, fmt.Sprintf("%v.", details))

			So(err.Info, ShouldResemble, &errorAVA.CallInfo{
				PackageName: "github.com/ver13/ava/test/common/error/general_test",
				FileName:    "deepCopyWrong_test.go",
				FuncName:    "TestDeepCopyWrong",
				Line:        50,
			})
		})
		Convey("Went all functions are OK ", func() {
			details := "Details"
			err := generalErrorAVA.DeepCopyWrong(nil, details)

			So(err, ShouldNotBeNil)

			str := err.String()
			json, errorJSON := err.ToJSON()

			So(str, ShouldResemble, "Group: General, Subgroup: General, Details: Details., Error: <nil>, Info: &{github.com/ver13/ava/test/common/error/general_test deepCopyWrong_test.go TestDeepCopyWrong 72}")
			So(errorJSON, ShouldBeNil)
			So(json, ShouldResemble, "{\n    \"Group\": 0,\n    \"Subgroup\": 0,\n    \"Err\": null,\n    \"Code\": 3,\n    \"Message\": \"Deep copy wrong.\",\n    \"Details\": \"Details.\",\n    \"Info\": {\n        \"PackageName\": \"github.com/ver13/ava/test/common/error/general_test\",\n        \"FileName\": \"deepCopyWrong_test.go\",\n        \"FuncName\": \"TestDeepCopyWrong\",\n        \"Line\": 72\n    }\n}")
		})
		Convey("Went it's wrong ", func() {
			details := "Details"
			err := generalErrorAVA.DeepCopyWrong(nil, details)

			So(err, ShouldNotBeNil)
			So(err.Details, ShouldNotResemble, details)
		})
	})
}
