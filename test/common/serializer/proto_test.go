package serializer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type postgresProto struct {
	User     string `proto:"user"`
	Password string `proto:"password"`
	Database string `proto:"db" commented:"true" comment:"not used anymore"`
}

type configProto struct {
	Postgres postgresProto `proto:"postgresYaml" comment:"Postgres configurationServiceI"`
}

type protoSuite struct {
	suite.Suite

	structData configProto

	serializer serializerAVA.SerializerI
}

func TestProtoInit(t *testing.T) {
	suite.Run(t, new(protoSuite))
}

func (r *protoSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *protoSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *protoSuite) SetupSuite() {
	r.T().Log("SetupSuite")
	r.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeProto)
}

func (r *protoSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *protoSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *protoSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *protoSuite) TestPROTO() {
	Convey("Given a struct ", r.T(), func() {
		Convey("When marshal to Pretty PROTO ", func() {
			r.structData = configProto{postgresProto{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "GmfDataBase"}}

			data, err := r.serializer.Serializer(r.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "{\n    \"postgresYaml\": {\n        \"user\": \"Ulises\",\n        \"password\": \"Z;Z@pZz9G)MFAw[5\",\n        \"db\": \"GmfDataBase\"\n    }\n}")
		})
	})
	Convey("Given a PROTO ", r.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			r.structData = configProto{}
			err := r.serializer.Deserializer([]byte("{\"postgresYaml\":{\"user\":\"Ulises\",\"password\":\"Z;Z@pZz9G)MFAw[5\",\"db\":\"GmfDataBase\"}}"), &r.structData)
			So(err, ShouldBeNil)
			So(r.structData, ShouldNotBeEmpty)
			So(r.structData.Postgres.User, ShouldResemble, "Ulises")
			So(r.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(r.structData.Postgres.Database, ShouldResemble, "GmfDataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			r.structData = configProto{}
			err := r.serializer.Deserializer([]byte("{\"postgres\":{\"user\":\"Ulises\",\"password\":\"Z;Z@pZz9G)MFAw[5\",\"db\":\"GmfDataBase\"}"), &r.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (r *protoSuite) TestPROTO_Serializer() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *protoSuite) TestPROTO_Deserializer() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *protoSuite) TestPROTO_String() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *protoSuite) TestPROTO_Type() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
