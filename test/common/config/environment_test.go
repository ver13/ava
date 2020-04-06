package config_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type environmentSuite struct {
	suite.Suite
}

func TestEnvironmentInit(t *testing.T) {
	suite.Run(t, new(environmentSuite))
}

func (r *environmentSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *environmentSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *environmentSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *environmentSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *environmentSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *environmentSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *environmentSuite) TestEnvironmen_Parser() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *environmentSuite) TestEnvironmen_Serializer() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
