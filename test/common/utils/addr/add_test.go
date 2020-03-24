package addr_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type addrAVASuite struct {
	suite.Suite
}

func TestAddrAVAInit(t *testing.T) {
	suite.Run(t, new(addrAVASuite))
}

func (a *addrAVASuite) BeforeTest() {
	a.T().Log("BeforeTest")
}

func (a *addrAVASuite) AfterTest() {
	a.T().Log("AfterTest")
}

func (a *addrAVASuite) SetupSuite() {
	a.T().Log("SetupSuite")
}

func (a *addrAVASuite) SetupTest() {
	a.T().Log("SetupTest")
}

func (a *addrAVASuite) TearDownSuite() {
	a.T().Log("TearDownSuite")
}

func (a *addrAVASuite) TearDownTest() {
	a.T().Log("TearDownTest")
}

func (a *addrAVASuite) Test_GetIPWithPrefix() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrAVASuite) Test_GetIP() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrAVASuite) Test_ResolveIPFromHostsFile() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrAVASuite) Test_IPs() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrAVASuite) Test_Extract() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrAVASuite) Test_IsPrivateIP() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
