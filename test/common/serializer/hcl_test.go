package serializer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type postgresHcl struct {
	User     string `hcl:"user"`
	Password string `hcl:"password"`
	Database string `hcl:"db" commented:"true" comment:"not used anymore"`
}

type configHcl struct {
	Postgres postgresHcl `hcl:"postgres" comment:"Postgres configurationServiceI"`
}

type hclSuite struct {
	suite.Suite

	structData configHcl

	serializer serializerAVA.SerializerI
}

func TestHclInit(t *testing.T) {
	suite.Run(t, new(hclSuite))
}

func (s *hclSuite) BeforeTest() {
	s.T().Log("BeforeTest")
}

func (s *hclSuite) AfterTest() {
	s.T().Log("AfterTest")
}

func (s *hclSuite) SetupSuite() {
	s.T().Log("SetupSuite")
	s.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeHcl)
}

func (s *hclSuite) SetupTest() {
	s.T().Log("SetupTest")
}

func (s *hclSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
}

func (s *hclSuite) TearDownTest() {
	s.T().Log("TearDownTest")
}

func (s *hclSuite) TestHCL() {
	Convey("Given a struct ", s.T(), func() {
		Convey("When marshal to HCL ", func() {
			s.structData = configHcl{postgresHcl{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "AVADataBase"}}

			data, err := s.serializer.Serializer(s.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "{\n    \"Postgres\": {\n        \"User\": \"Ulises\",\n        \"Password\": \"Z;Z@pZz9G)MFAw[5\",\n        \"Database\": \"AVADataBase\"\n    }\n}")
		})
	})
	Convey("Given a HCL ", s.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			s.structData = configHcl{}
			err := s.serializer.Deserializer([]byte("{\"postgres\":{\"user\":\"Ulises\",\"password\":\"Z;Z@pZz9G)MFAw[5\",\"db\":\"AVADataBase\"}}"), &s.structData)
			So(err, ShouldBeNil)
			So(s.structData, ShouldNotBeEmpty)
			So(s.structData.Postgres.User, ShouldResemble, "Ulises")
			So(s.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(s.structData.Postgres.Database, ShouldResemble, "AVADataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			s.structData = configHcl{}
			err := s.serializer.Deserializer([]byte("jjk"), &s.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (s *hclSuite) TestHCL_Serializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *hclSuite) TestHCL_Deserializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *hclSuite) TestHCL_String() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *hclSuite) TestHCL_Type() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
