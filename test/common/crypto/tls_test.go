package crypto_test

import (
	"testing"
	
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type tlsSuite struct {
	suite.Suite
}

func TestTLSInit(t *testing.T) {
	suite.Run(t, new(tlsSuite))
}

func (tls *tlsSuite) BeforeTest() {
	tls.T().Log("BeforeTest")
}

func (tls *tlsSuite) AfterTest() {
	tls.T().Log("AfterTest")
}

func (tls *tlsSuite) SetupSuite() {
	tls.T().Log("SetupSuite")
}

func (tls *tlsSuite) SetupTest() {
	tls.T().Log("SetupTest")
}

func (tls *tlsSuite) TearDownSuite() {
	tls.T().Log("TearDownSuite")
}

func (tls *tlsSuite) TearDownTest() {
	tls.T().Log("TearDownTest")
}

func (tls *tlsSuite) Test_Certificate() {
	Convey("Given a ", tls.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
