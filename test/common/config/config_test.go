package config_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type configSuite struct {
	suite.Suite
}

func TestConfigInit(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (r *configSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *configSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *configSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *configSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *configSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *configSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *configSuite) Test_Logger() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_Backend() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_TLS() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_CORS() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_Endpoint() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_API() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_Environment() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *configSuite) Test_Configuration() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
