package serializer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type postgresYaml struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"db" commented:"true" comment:"not used anymore"`
}
type configYaml struct {
	Postgres postgresYaml `yaml:"postgresYaml" comment:"Postgres configurationServiceI"`
}

type yamlSuite struct {
	suite.Suite

	structData configYaml

	serializer serializerAVA.SerializerI
}

func TestYamlInit(t *testing.T) {
	suite.Run(t, new(yamlSuite))
}

func (r *yamlSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *yamlSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *yamlSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *yamlSuite) SetupTest() {
	r.T().Log("SetupTest")

	r.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeYaml)
}

func (r *yamlSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *yamlSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *yamlSuite) TestYAML() {
	Convey("Given a struct ", r.T(), func() {
		Convey("When marshal to YAML ", func() {
			r.structData = configYaml{postgresYaml{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "GmfDataBase"}}

			data, err := r.serializer.Serializer(r.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "postgresYaml:\n  user: Ulises\n  password: Z;Z@pZz9G)MFAw[5\n  db: GmfDataBase\n")
		})
	})
	Convey("Given a YAML ", r.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			r.structData = configYaml{}
			err := r.serializer.Deserializer([]byte("postgresYaml:\n  user: Ulises\n  password: Z;Z@pZz9G)MFAw[5\n  db: GmfDataBase\n"), &r.structData)
			So(err, ShouldBeNil)
			So(r.structData, ShouldNotBeEmpty)
			So(r.structData.Postgres.User, ShouldResemble, "Ulises")
			So(r.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(r.structData.Postgres.Database, ShouldResemble, "GmfDataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			r.structData = configYaml{}
			err := r.serializer.Deserializer([]byte("postgresYaml:  user: Ulises  password: Z;Z@pZz9G)MFAw[5  db: GmfDataBase"), &r.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (s *yamlSuite) TestYAML_Serializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *yamlSuite) TestYAML_Deserializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *yamlSuite) TestYAML_String() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *yamlSuite) TestYAML_Type() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
