package config_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type loggerSuite struct {
	suite.Suite
}

func TestLoggerInit(t *testing.T) {
	suite.Run(t, new(loggerSuite))
}

func (r *loggerSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *loggerSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *loggerSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *loggerSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *loggerSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *loggerSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *loggerSuite) TestLogger_Parser() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *loggerSuite) TestLogger_Serializer() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
