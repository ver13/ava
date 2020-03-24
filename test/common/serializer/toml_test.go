package serializer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type postgresToml struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"db" commented:"true" comment:"not used anymore"`
}

type configToml struct {
	Postgres postgresToml `toml:"postgresYaml" comment:"Postgres configurationServiceI"`
}

type tomlSuite struct {
	suite.Suite

	structData configToml

	serializer serializerAVA.SerializerI
}

func TestTomlInit(t *testing.T) {
	suite.Run(t, new(tomlSuite))
}

func (r *tomlSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *tomlSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *tomlSuite) SetupSuite() {
	r.T().Log("SetupSuite")

	r.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeToml)
}

func (r *tomlSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *tomlSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *tomlSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *tomlSuite) TestTOML() {
	Convey("Given a struct ", r.T(), func() {
		Convey("When marshal to TOML ", func() {
			r.structData = configToml{postgresToml{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "AVADataBase"}}

			data, err := r.serializer.Serializer(r.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "[postgresYaml]\n  user = \"Ulises\"\n  password = \"Z;Z@pZz9G)MFAw[5\"\n  db = \"AVADataBase\"\n")
		})
	})
	Convey("Given a TOML ", r.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			r.structData = configToml{}
			err := r.serializer.Deserializer([]byte("\n# Postgres configurationServiceI\n[postgresYaml]\n  db = \"AVADataBase\"\n  password = \"Z;Z@pZz9G)MFAw[5\"\n  user = \"Ulises\"\n"), &r.structData)
			So(err, ShouldBeNil)
			So(r.structData, ShouldNotBeEmpty)
			So(r.structData.Postgres.User, ShouldResemble, "Ulises")
			So(r.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(r.structData.Postgres.Database, ShouldResemble, "AVADataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			r.structData = configToml{}
			err := r.serializer.Deserializer([]byte("Postgres configurationServiceI\n[postgres]\n\n  # not used anymore\n  db = \"AVADataBase\"\n  password = \"Z;Z@pZz9G)MFAw[5\"\n  user = \"Ulises\"\n"), &r.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (s *tomlSuite) TestTOML_Serializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *tomlSuite) TestTOML_Deserializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *tomlSuite) TestTOML_String() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *tomlSuite) TestTOML_Type() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
