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
}

func TestHclInit(t *testing.T) {
	suite.Run(t, new(hclSuite))
}

func (h *hclSuite) BeforeTest() {
	h.T().Log("BeforeTest")
}

func (h *hclSuite) AfterTest() {
	h.T().Log("AfterTest")
}

func (h *hclSuite) SetupSuite() {
	h.T().Log("SetupSuite")
}

func (h *hclSuite) SetupTest() {
	h.T().Log("SetupTest")
}

func (h *hclSuite) TearDownSuite() {
	h.T().Log("TearDownSuite")
}

func (h *hclSuite) TearDownTest() {
	h.T().Log("TearDownTest")
}

func (h *hclSuite) TestHCL() {
	Convey("Given a struct ", h.T(), func() {
		Convey("When marshal to HCL ", func() {
			h.structData = configHcl{postgresHcl{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "AVADataBase"}}
			s := serializerAVA.GetSerializer(serializerAVA.SerializerTypeHcl)

			data, err := s.Serializer(h.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "{\n    \"Postgres\": {\n        \"User\": \"Ulises\",\n        \"Password\": \"Z;Z@pZz9G)MFAw[5\",\n        \"Database\": \"AVADataBase\"\n    }\n}")
		})
	})
	Convey("Given a HCL ", h.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			h.structData = configHcl{}
			s := serializerAVA.GetSerializer(serializerAVA.SerializerTypeHcl)

			err := s.Deserializer([]byte("{\"postgres\":{\"user\":\"Ulises\",\"password\":\"Z;Z@pZz9G)MFAw[5\",\"db\":\"AVADataBase\"}}"), &h.structData)
			So(err, ShouldBeNil)
			So(h.structData, ShouldNotBeEmpty)
			So(h.structData.Postgres.User, ShouldResemble, "Ulises")
			So(h.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(h.structData.Postgres.Database, ShouldResemble, "AVADataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			h.structData = configHcl{}
			s := serializerAVA.GetSerializer(serializerAVA.SerializerTypeHcl)

			err := s.Deserializer([]byte("jjk"), &h.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (h *hclSuite) TestHCL_Serializer() {
	Convey("Given a ", h.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (h *hclSuite) TestHCL_Deserializer() {
	Convey("Given a ", h.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (h *hclSuite) TestHCL_String() {
	Convey("Given a ", h.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (h *hclSuite) TestHCL_Type() {
	Convey("Given a ", h.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
