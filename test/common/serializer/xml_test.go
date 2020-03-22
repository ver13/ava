package serializer_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type postgresXml struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
	Database string `xml:"db" commented:"true" comment:"not used anymore"`
}
type configXml struct {
	Postgres postgresXml `xml:"postgresYaml" comment:"Postgres configurationServiceI"`
}

type xmlSuite struct {
	suite.Suite

	structData configXml

	serializer serializerAVA.SerializerI
}

func TestXmlInit(t *testing.T) {
	suite.Run(t, new(xmlSuite))
}

func (r *xmlSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *xmlSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *xmlSuite) SetupSuite() {
	r.T().Log("SetupSuite")

	r.serializer = serializerAVA.GetSerializer(serializerAVA.SerializerTypeXml)
}

func (r *xmlSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *xmlSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *xmlSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *xmlSuite) TestXML() {
	Convey("Given a struct ", r.T(), func() {
		Convey("When marshal to Pretty XML ", func() {
			r.structData = configXml{postgresXml{User: "Ulises", Password: "Z;Z@pZz9G)MFAw[5", Database: "GmfDataBase"}}

			data, err := r.serializer.Serializer(r.structData)
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
			So(string(data), ShouldResemble, "<configXml>\n    <postgresYaml>\n        <user>Ulises</user>\n        <password>Z;Z@pZz9G)MFAw[5</password>\n        <db>GmfDataBase</db>\n    </postgresYaml>\n</configXml>")
		})
	})
	Convey("Given a XML ", r.T(), func() {
		Convey("When unmarshal to struct is OK", func() {
			r.structData = configXml{}
			err := r.serializer.Deserializer([]byte("<configXml><postgresYaml><user>Ulises</user><password>Z;Z@pZz9G)MFAw[5</password><db>GmfDataBase</db></postgresYaml></configXml>"), &r.structData)
			So(err, ShouldBeNil)
			So(r.structData, ShouldNotBeEmpty)
			So(r.structData.Postgres.User, ShouldResemble, "Ulises")
			So(r.structData.Postgres.Password, ShouldResemble, "Z;Z@pZz9G)MFAw[5")
			So(r.structData.Postgres.Database, ShouldResemble, "GmfDataBase")
		})
		Convey("When unmarshal to struct is Failed", func() {
			r.structData = configXml{}
			err := r.serializer.Deserializer([]byte("<error><postgres><user>Ulises</user><password>Z;Z@pZz9G)MFAw[5</password><db>GmfDataBase</db></postgresYaml></configXml>"), &r.structData)
			So(err, ShouldNotBeNil)
		})
	})
}

func (s *xmlSuite) TestXML_Serializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *xmlSuite) TestXML_Deserializer() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *xmlSuite) TestXML_String() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *xmlSuite) TestXML_Type() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
