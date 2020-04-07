package version_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	versionAVA "github.com/ver13/ava/pkg/common/version"
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
		tests := []struct {
			version string
			err     bool
		}{
			{"1.2.3", false},
			{"v1.2.3", false},
			{"1.0", false},
			{"v1.0", false},
			{"1", false},
			{"v1", false},
			{"1.2.beta", true},
			{"v1.2.beta", true},
			{"foo", true},
			{"1.2-5", false},
			{"v1.2-5", false},
			{"1.2-beta.5", false},
			{"v1.2-beta.5", false},
			{"\n1.2", true},
			{"\nv1.2", true},
			{"1.2.0-x.Y.0+metadata", false},
			{"v1.2.0-x.Y.0+metadata", false},
			{"1.2.0-x.Y.0+metadata-width-hypen", false},
			{"v1.2.0-x.Y.0+metadata-width-hypen", false},
			{"1.2.3-rc1-with-hypen", false},
			{"v1.2.3-rc1-with-hypen", false},
			{"1.2.3.4", true},
			{"v1.2.3.4", true},
		}

		for _, tc := range tests {
			versionAVA.SemanticVersion = tc.version
			_, err := versionAVA.GetInstance()
			if tc.err && err != nil {
				r.T().Fatalf("expected error for version: %s", tc.version)
			} else if !tc.err && err != nil {
				r.T().Fatalf("error for version %s: %s", tc.version, err)
			}
		}
	})
}

func (r *versionInfoSuite) TestVersionInfo_Parts() {
	Convey("VersionInfo SemanticVersion Parts ", r.T(), func() {
		versionAVA.SemanticVersion = "1.2.3-beta.1+build.123"
		v, err := versionAVA.GetInstance()
		if err != nil {
			r.T().Error("Error parsing version 1.2.3-beta.1+build.123")
		}

		if v.Major() != 1 {
			r.T().Error("Major() returning wrong value")
		}
		if v.Minor() != 2 {
			r.T().Error("Minor() returning wrong value")
		}
		if v.Patch() != 3 {
			r.T().Error("Patch() returning wrong value")
		}
		if v.Prerelease() != "beta.1" {
			r.T().Error("Prerelease() returning wrong value")
		}
		if v.Metadata() != "build.123" {
			r.T().Error("Metadata() returning wrong value")
		}
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
