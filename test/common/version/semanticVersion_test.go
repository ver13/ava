package version_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type semanticVersionSuite struct {
	suite.Suite
}

func TestSemanticVersionInit(t *testing.T) {
	suite.Run(t, new(semanticVersionSuite))
}

func (r *semanticVersionSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *semanticVersionSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *semanticVersionSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *semanticVersionSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *semanticVersionSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *semanticVersionSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (s *semanticVersionSuite) TestTestVersionInfo_String() {
	Convey("Semantic version string ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_Equals() {
	Convey("Semantic version equals ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_EQ() {
	Convey("Semantic version EQ ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_NE() {
	Convey("Semantic version NE ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_GT() {
	Convey("Semantic version GT ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_GTE() {
	Convey("Semantic version GTE ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_GE() {
	Convey("Semantic version GE ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_LT() {
	Convey("Semantic version LT ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_LTE() {
	Convey("Semantic version LTE ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_LE() {
	Convey("Semantic version LE ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_IncrementPatch() {
	Convey("Semantic version increment patch ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_IncrementMinor() {
	Convey("Semantic version increment minor", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_IncrementMajor() {
	Convey("Semantic version increment major ", s.T(), func() {

	})
}

func (s *semanticVersionSuite) TestTestVersionInfo_Validate() {
	Convey("Semantic version validate ", s.T(), func() {

	})
}
