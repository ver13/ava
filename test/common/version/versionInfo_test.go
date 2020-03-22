package version_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type versionInfoSuite struct {
	suite.Suite
}

func TestVersionInit(t *testing.T) {
	suite.Run(t, new(versionInfoSuite))
}

func (r *versionInfoSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *versionInfoSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *versionInfoSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *versionInfoSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *versionInfoSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *versionInfoSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *versionInfoSuite) TestVersionInfo_SemanticVersion() {
	Convey("VersionInfo Semantic version ", r.T(), func() {

	})
}

func (r *versionInfoSuite) TestVersionInfo_Name() {
	Convey("VersionInfo Name ", r.T(), func() {

	})
}

func (r *versionInfoSuite) TestVersionInfo_ServerName() {
	Convey("VersionInfo Server name ", r.T(), func() {

	})
}

func (r *versionInfoSuite) TestVersionInfo_ClientName() {
	Convey("VersionInfo Client name ", r.T(), func() {

	})
}

func (r *versionInfoSuite) TestVersionInfo_GitCommit() {
	Convey("VersionInfo Git commit ", r.T(), func() {

	})
}

func (r *versionInfoSuite) TestVersionInfo_GoVersion() {
	Convey("VersionInfo Go version", r.T(), func() {

	})
}

func (r *versionInfoSuite) TestVersionInfo_String() {
	Convey("VersionInfo String ", r.T(), func() {

	})
}
