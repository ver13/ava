package time_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/time/timezones"
)

type timezonesSuite struct {
	suite.Suite
}

func TestTimezonesInit(t *testing.T) {
	suite.Run(t, new(timezonesSuite))
}

func (r *timezonesSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *timezonesSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *timezonesSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *timezonesSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *timezonesSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *timezonesSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *timezonesSuite) Test_TimezoneConfig() {
	Convey("Given a Timezones ", r.T(), func() {
		tz1 := New()

		So(tz1.GetSupported(), ShouldNotBeEmpty)
	})
}

func (r *timezonesSuite) TestDefaultUserTimezone() {
	Convey("Given a Timezones ", r.T(), func() {
		defaultTimezone := DefaultUserTimezone()

		So("true", ShouldResemble, defaultTimezone["useAutomaticTimezone"])
		So(defaultTimezone["automaticTimezone"], ShouldBeEmpty)
		So(defaultTimezone["manualTimezone"], ShouldBeEmpty)

		defaultTimezone["useAutomaticTimezone"] = "false"
		defaultTimezone["automaticTimezone"] = "EST"
		defaultTimezone["manualTimezone"] = "AST"

		defaultTimezone2 := DefaultUserTimezone()
		So("true", ShouldResemble, defaultTimezone2["useAutomaticTimezone"])
		So(defaultTimezone2["automaticTimezone"], ShouldBeEmpty)
		So(defaultTimezone2["manualTimezone"], ShouldBeEmpty)
	})
}
