package addr_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type addrGmfSuite struct {
	suite.Suite
}

func TestAddrGmfInit(t *testing.T) {
	suite.Run(t, new(addrGmfSuite))
}

func (a *addrGmfSuite) BeforeTest() {
	a.T().Log("BeforeTest")
}

func (a *addrGmfSuite) AfterTest() {
	a.T().Log("AfterTest")
}

func (a *addrGmfSuite) SetupSuite() {
	a.T().Log("SetupSuite")
}

func (a *addrGmfSuite) SetupTest() {
	a.T().Log("SetupTest")
}

func (a *addrGmfSuite) TearDownSuite() {
	a.T().Log("TearDownSuite")
}

func (a *addrGmfSuite) TearDownTest() {
	a.T().Log("TearDownTest")
}

func (a *addrGmfSuite) Test_GetIPWithPrefix() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrGmfSuite) Test_GetIP() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrGmfSuite) Test_ResolveIPFromHostsFile() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrGmfSuite) Test_IPs() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrGmfSuite) Test_Extract() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (a *addrGmfSuite) Test_IsPrivateIP() {
	Convey("Given a ", a.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
