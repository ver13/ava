package error_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type subgroupTypeSuite struct {
	suite.Suite

	subgroup errorAVA.Subgroup
}

func TestSubgroupTypeInit(t *testing.T) {
	suite.Run(t, new(subgroupTypeSuite))
}

func (r *subgroupTypeSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *subgroupTypeSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *subgroupTypeSuite) SetupSuite() {
	r.T().Log("SetupSuite")

	r.subgroup = errorAVA.SubgroupUnknown
}

func (r *subgroupTypeSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *subgroupTypeSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *subgroupTypeSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *subgroupTypeSuite) TestSubgroupType_ParseSubgroup() {
	Convey("Given a Subgroup type", r.T(), func() {
		Convey("Went its function ParseSubgroup() it's OK ", func() {
			code, err := errorAVA.ParseSubgroup("General")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupGeneral)

			code, err = errorAVA.ParseSubgroup("DiscoveryService")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupDiscoveryService)

			code, err = errorAVA.ParseSubgroup("BrokerService")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupBrokerService)

			code, err = errorAVA.ParseSubgroup("CircuitBreakerService")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupCircuitBreakerService)

			code, err = errorAVA.ParseSubgroup("MetricsService")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupMetricsService)

			code, err = errorAVA.ParseSubgroup("Client")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupClient)

			code, err = errorAVA.ParseSubgroup("Server")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupServer)

			code, err = errorAVA.ParseSubgroup("Selected")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupSelected)

			code, err = errorAVA.ParseSubgroup("Unknown")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupUnknown)
		})
	})
}

func (r *subgroupTypeSuite) TestSubgroupType_String() {
	Convey("Given a Subgroup type", r.T(), func() {
		Convey("Went its function String() it's OK ", func() {
			So(errorAVA.SubgroupGeneral.String(), ShouldEqual, "General")
			So(errorAVA.SubgroupDiscoveryService.String(), ShouldEqual, "DiscoveryService")
			So(errorAVA.SubgroupBrokerService.String(), ShouldEqual, "BrokerService")
			So(errorAVA.SubgroupCircuitBreakerService.String(), ShouldEqual, "CircuitBreakerService")
			So(errorAVA.SubgroupMetricsService.String(), ShouldEqual, "MetricsService")
			So(errorAVA.SubgroupClient.String(), ShouldEqual, "Client")
			So(errorAVA.SubgroupServer.String(), ShouldEqual, "Server")
			So(errorAVA.SubgroupSelected.String(), ShouldEqual, "Selected")
			So(errorAVA.SubgroupUnknown.String(), ShouldEqual, "Unknown")
		})
	})
}

func (r *subgroupTypeSuite) TestSubgroupType_MarshalText() {
	Convey("Given a Subgroup type", r.T(), func() {
		Convey("Went its function String() it's OK ", func() {
			b, _ := errorAVA.SubgroupGeneral.MarshalText()
			So(string(b), ShouldEqual, "General")

			b, _ = errorAVA.SubgroupDiscoveryService.MarshalText()
			So(string(b), ShouldEqual, "DiscoveryService")

			b, _ = errorAVA.SubgroupBrokerService.MarshalText()
			So(string(b), ShouldEqual, "BrokerService")

			b, _ = errorAVA.SubgroupCircuitBreakerService.MarshalText()
			So(string(b), ShouldEqual, "CircuitBreakerService")

			b, _ = errorAVA.SubgroupMetricsService.MarshalText()
			So(string(b), ShouldEqual, "MetricsService")

			b, _ = errorAVA.SubgroupClient.MarshalText()
			So(string(b), ShouldEqual, "Client")

			b, _ = errorAVA.SubgroupServer.MarshalText()
			So(string(b), ShouldEqual, "Server")

			b, _ = errorAVA.SubgroupSelected.MarshalText()
			So(string(b), ShouldEqual, "Selected")

			b, _ = errorAVA.SubgroupUnknown.MarshalText()
			So(string(b), ShouldEqual, "Subgroup(17)")
		})
	})
}

func (r *subgroupTypeSuite) TestSubgroupType_UnmarshalText() {
	Convey("Given a Subgroup type", r.T(), func() {
		Convey("Went its function String() it's OK ", func() {
			code, err := r.subgroup.UnmarshalText([]byte("General"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupGeneral)

			code, err = r.subgroup.UnmarshalText([]byte("DiscoveryService"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupDiscoveryService)

			code, err = r.subgroup.UnmarshalText([]byte("BrokerService"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupBrokerService)

			code, err = r.subgroup.UnmarshalText([]byte("CircuitBreakerService"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupCircuitBreakerService)

			code, err = r.subgroup.UnmarshalText([]byte("MetricsService"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupMetricsService)

			code, err = r.subgroup.UnmarshalText([]byte("Client"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupClient)

			code, err = r.subgroup.UnmarshalText([]byte("Server"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupServer)

			code, err = r.subgroup.UnmarshalText([]byte("Selected"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupSelected)

			code, err = r.subgroup.UnmarshalText([]byte("Unknown"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.SubgroupUnknown)
		})
		Convey("Went its function String() it's failure ", func() {
			code, err := r.subgroup.UnmarshalText([]byte("Failure"))
			So(err, ShouldNotBeNil)
			So(err, ShouldHaveSameTypeAs, errorAVA.SubgroupTypeUnknownSkip(nil, fmt.Sprintf("%s is not a valid Group", "Failure"), 4))
			So(code, ShouldEqual, errorAVA.Group(0))
		})
	})
}
