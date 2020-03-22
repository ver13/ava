package serializer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type postgresJson struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"db" commented:"true" comment:"not used anymore"`
}

type configJson struct {
	Postgres postgresJson `json:"postgresYaml" comment:"Postgres configurationServiceI"`
}

type jsonSuite struct {
	suite.Suite

	structData configJson

	serializer serializerAVA.SerializerI
}

func TestJsonInit(t *testing.T) {
	suite.Run(t, new(jsonSuite))
}

func (r *jsonSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *jsonSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *jsonSuite) SetupSuite() {
	r.T().Log("SetupSuite")
	r.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeJson)
}

func (r *jsonSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *jsonSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *jsonSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *jsonSuite) TestJSON() {
	Convey("Given a struct ", r.T(), func() {
		Convey("When marshal to Pretty JSON ", func() {
			r.structData = configJson{postgresJson{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "GmfDataBase"}}

			data, err := r.serializer.Serializer(r.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "{\n    \"postgresYaml\": {\n        \"user\": \"Ulises\",\n        \"password\": \"Z;Z@pZz9G)MFAw[5\",\n        \"db\": \"GmfDataBase\"\n    }\n}")
		})
	})
	Convey("Given a JSON ", r.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			r.structData = configJson{}
			err := r.serializer.Deserializer([]byte("{\"postgresYaml\":{\"user\":\"Ulises\",\"password\":\"Z;Z@pZz9G)MFAw[5\",\"db\":\"GmfDataBase\"}}"), &r.structData)
			So(err, ShouldBeNil)
			So(r.structData, ShouldNotBeEmpty)
			So(r.structData.Postgres.User, ShouldResemble, "Ulises")
			So(r.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(r.structData.Postgres.Database, ShouldResemble, "GmfDataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			r.structData = configJson{}
			err := r.serializer.Deserializer([]byte("{\"postgres\":{\"user\":\"Ulises\",\"password\":\"Z;Z@pZz9G)MFAw[5\",\"db\":\"GmfDataBase\"}"), &r.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (s *jsonSuite) TestJSON_Serializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *jsonSuite) TestJSON_Deserializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *jsonSuite) TestJSON_String() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *jsonSuite) TestJSON_Type() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
