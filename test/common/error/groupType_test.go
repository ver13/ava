package error_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type groupTypeSuite struct {
	suite.Suite

	group errorAVA.Group
}

func TestGroupTypeInit(t *testing.T) {
	suite.Run(t, new(groupTypeSuite))
}

func (r *groupTypeSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *groupTypeSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *groupTypeSuite) SetupSuite() {
	r.T().Log("SetupSuite")

	r.group = errorAVA.GroupUnknown
}

func (r *groupTypeSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *groupTypeSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *groupTypeSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *groupTypeSuite) TestGroupType_ParseGroup() {
	Convey("Given a Group type", r.T(), func() {
		Convey("Went its function ParseGroup() it's OK ", func() {
			code, err := errorAVA.ParseGroup("General")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupGeneral)

			code, err = errorAVA.ParseGroup("Model")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupModel)

			code, err = errorAVA.ParseGroup("Serializer")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupSerializer)

			code, err = errorAVA.ParseGroup("Encoder")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupEncoder)

			code, err = errorAVA.ParseGroup("Server")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupServer)

			code, err = errorAVA.ParseGroup("Config")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupConfig)

			code, err = errorAVA.ParseGroup("Logger")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupLogger)

			code, err = errorAVA.ParseGroup("File")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupFile)

			code, err = errorAVA.ParseGroup("Blockchain")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupBlockchain)

			code, err = errorAVA.ParseGroup("Database")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupDatabase)

			code, err = errorAVA.ParseGroup("Http")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupHttp)

			code, err = errorAVA.ParseGroup("Microservice")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupMicroservice)

			code, err = errorAVA.ParseGroup("MessageCoder")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupMessageCoder)

			code, err = errorAVA.ParseGroup("Time")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupTime)

			code, err = errorAVA.ParseGroup("ApiTime")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupApiTime)

			code, err = errorAVA.ParseGroup("Transport")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupTransport)

			code, err = errorAVA.ParseGroup("Compress")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupCompress)

			code, err = errorAVA.ParseGroup("IO")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupIO)

			code, err = errorAVA.ParseGroup("Crypto")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupCrypto)

			code, err = errorAVA.ParseGroup("QR")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupQR)

			code, err = errorAVA.ParseGroup("Validator")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupValidator)

			code, err = errorAVA.ParseGroup("String")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupString)

			code, err = errorAVA.ParseGroup("Utils")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupUtils)

			code, err = errorAVA.ParseGroup("Client")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupClient)

			code, err = errorAVA.ParseGroup("GeneratorEnum")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupGeneratorEnum)

			code, err = errorAVA.ParseGroup("Router")
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupRouter)
		})
		Convey("Went its function String() it's failure ", func() {
			code, err := errorAVA.ParseGroup("Failure")
			So(err, ShouldNotBeNil)
			So(err, ShouldHaveSameTypeAs, errorAVA.GroupTypeUnknownSkip(nil, fmt.Sprintf("%s is not a valid Group", "Failure"), 4))
			So(code, ShouldEqual, errorAVA.GroupUnknown)
		})
	})
}

func (r *groupTypeSuite) TestGroupType_String() {
	Convey("Given a Group type", r.T(), func() {
		Convey("Went its function String() it's OK ", func() {
			So(errorAVA.GroupGeneral.String(), ShouldEqual, "General")
			So(errorAVA.GroupModel.String(), ShouldEqual, "Model")
			So(errorAVA.GroupSerializer.String(), ShouldEqual, "Serializer")
			So(errorAVA.GroupEncoder.String(), ShouldEqual, "Encoder")
			So(errorAVA.GroupServer.String(), ShouldEqual, "Server")
			So(errorAVA.GroupConfig.String(), ShouldEqual, "Config")
			So(errorAVA.GroupLogger.String(), ShouldEqual, "Logger")
			So(errorAVA.GroupFile.String(), ShouldEqual, "File")
			So(errorAVA.GroupBlockchain.String(), ShouldEqual, "Blockchain")
			So(errorAVA.GroupDatabase.String(), ShouldEqual, "Database")
			So(errorAVA.GroupHttp.String(), ShouldEqual, "Http")
			So(errorAVA.GroupMicroservice.String(), ShouldEqual, "Microservice")
			So(errorAVA.GroupMessageCoder.String(), ShouldEqual, "MessageCoder")
			So(errorAVA.GroupTime.String(), ShouldEqual, "Time")
			So(errorAVA.GroupApiTime.String(), ShouldEqual, "ApiTime")
			So(errorAVA.GroupTransport.String(), ShouldEqual, "Transport")
			So(errorAVA.GroupCompress.String(), ShouldEqual, "Compress")
			So(errorAVA.GroupIO.String(), ShouldEqual, "IO")
			So(errorAVA.GroupCrypto.String(), ShouldEqual, "Crypto")
			So(errorAVA.GroupQR.String(), ShouldEqual, "QR")
			So(errorAVA.GroupValidator.String(), ShouldEqual, "Validator")
			So(errorAVA.GroupString.String(), ShouldEqual, "String")
			So(errorAVA.GroupUtils.String(), ShouldEqual, "Utils")
			So(errorAVA.GroupClient.String(), ShouldEqual, "Client")
			So(errorAVA.GroupGeneratorEnum.String(), ShouldEqual, "GeneratorEnum")
			So(errorAVA.GroupRouter.String(), ShouldEqual, "Router")
		})
	})
}

func (r *groupTypeSuite) TestGroupType_MarshalText() {
	Convey("Given a Group type", r.T(), func() {
		Convey("Went its function String() it's OK ", func() {
			b, _ := errorAVA.GroupGeneral.MarshalText()
			So(string(b), ShouldEqual, "General")

			b, _ = errorAVA.GroupModel.MarshalText()
			So(string(b), ShouldEqual, "Model")

			b, _ = errorAVA.GroupSerializer.MarshalText()
			So(string(b), ShouldEqual, "Serializer")

			b, _ = errorAVA.GroupEncoder.MarshalText()
			So(string(b), ShouldEqual, "Encoder")

			b, _ = errorAVA.GroupServer.MarshalText()
			So(string(b), ShouldEqual, "Server")

			b, _ = errorAVA.GroupConfig.MarshalText()
			So(string(b), ShouldEqual, "Config")

			b, _ = errorAVA.GroupLogger.MarshalText()
			So(string(b), ShouldEqual, "Logger")

			b, _ = errorAVA.GroupFile.MarshalText()
			So(string(b), ShouldEqual, "File")

			b, _ = errorAVA.GroupBlockchain.MarshalText()
			So(string(b), ShouldEqual, "Blockchain")

			b, _ = errorAVA.GroupDatabase.MarshalText()
			So(string(b), ShouldEqual, "Database")

			b, _ = errorAVA.GroupHttp.MarshalText()
			So(string(b), ShouldEqual, "Http")

			b, _ = errorAVA.GroupMicroservice.MarshalText()
			So(string(b), ShouldEqual, "Microservice")

			b, _ = errorAVA.GroupMessageCoder.MarshalText()
			So(string(b), ShouldEqual, "MessageCoder")

			b, _ = errorAVA.GroupTime.MarshalText()
			So(string(b), ShouldEqual, "Time")

			b, _ = errorAVA.GroupApiTime.MarshalText()
			So(string(b), ShouldEqual, "ApiTime")

			b, _ = errorAVA.GroupTransport.MarshalText()
			So(string(b), ShouldEqual, "Transport")

			b, _ = errorAVA.GroupCompress.MarshalText()
			So(string(b), ShouldEqual, "Compress")

			b, _ = errorAVA.GroupIO.MarshalText()
			So(string(b), ShouldEqual, "IO")

			b, _ = errorAVA.GroupCrypto.MarshalText()
			So(string(b), ShouldEqual, "Crypto")

			b, _ = errorAVA.GroupQR.MarshalText()
			So(string(b), ShouldEqual, "QR")

			b, _ = errorAVA.GroupValidator.MarshalText()
			So(string(b), ShouldEqual, "Validator")

			b, _ = errorAVA.GroupString.MarshalText()
			So(string(b), ShouldEqual, "String")

			b, _ = errorAVA.GroupUtils.MarshalText()
			So(string(b), ShouldEqual, "Utils")

			b, _ = errorAVA.GroupClient.MarshalText()
			So(string(b), ShouldEqual, "Client")

			b, _ = errorAVA.GroupGeneratorEnum.MarshalText()
			So(string(b), ShouldEqual, "GeneratorEnum")

			b, _ = errorAVA.GroupRouter.MarshalText()
			So(string(b), ShouldEqual, "Router")
		})
	})
}

func (r *groupTypeSuite) TestGroupType_UnmarshalText() {
	Convey("Given a Group type", r.T(), func() {
		Convey("Went its function String() it's OK ", func() {
			code, err := r.group.UnmarshalText([]byte("General"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupGeneral)

			code, err = r.group.UnmarshalText([]byte("Model"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupModel)

			code, err = r.group.UnmarshalText([]byte("Serializer"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupSerializer)

			code, err = r.group.UnmarshalText([]byte("Encoder"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupEncoder)

			code, err = r.group.UnmarshalText([]byte("Server"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupServer)

			code, err = r.group.UnmarshalText([]byte("Config"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupConfig)

			code, err = r.group.UnmarshalText([]byte("Logger"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupLogger)

			code, err = r.group.UnmarshalText([]byte("File"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupFile)

			code, err = r.group.UnmarshalText([]byte("Blockchain"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupBlockchain)

			code, err = r.group.UnmarshalText([]byte("Database"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupDatabase)

			code, err = r.group.UnmarshalText([]byte("Http"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupHttp)

			code, err = r.group.UnmarshalText([]byte("Microservice"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupMicroservice)

			code, err = r.group.UnmarshalText([]byte("MessageCoder"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupMessageCoder)

			code, err = r.group.UnmarshalText([]byte("Time"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupTime)

			code, err = r.group.UnmarshalText([]byte("ApiTime"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupApiTime)

			code, err = r.group.UnmarshalText([]byte("Transport"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupTransport)

			code, err = r.group.UnmarshalText([]byte("Compress"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupCompress)

			code, err = r.group.UnmarshalText([]byte("IO"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupIO)

			code, err = r.group.UnmarshalText([]byte("Crypto"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupCrypto)

			code, err = r.group.UnmarshalText([]byte("QR"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupQR)

			code, err = r.group.UnmarshalText([]byte("Validator"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupValidator)

			code, err = r.group.UnmarshalText([]byte("String"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupString)

			code, err = r.group.UnmarshalText([]byte("Utils"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupUtils)

			code, err = r.group.UnmarshalText([]byte("Client"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupClient)

			code, err = r.group.UnmarshalText([]byte("GeneratorEnum"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupGeneratorEnum)

			code, err = r.group.UnmarshalText([]byte("Router"))
			So(err, ShouldBeNil)
			So(code, ShouldEqual, errorAVA.GroupRouter)
		})
		Convey("Went its function String() it's failure ", func() {
			code, err := r.group.UnmarshalText([]byte("Failure"))
			So(err, ShouldNotBeNil)
			So(err, ShouldHaveSameTypeAs, errorAVA.GroupTypeUnknownSkip(nil, fmt.Sprintf("%s is not a valid Group", "Failure"), 4))
			So(code, ShouldEqual, errorAVA.GroupUnknown)
		})
	})
}
