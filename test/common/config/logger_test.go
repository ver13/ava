package config_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type loggerSuite struct {
	suite.Suite

	serializer *serializerAVA.Serializer
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
	r.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeJson)
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
