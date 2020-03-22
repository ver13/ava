package error_test

import (
	"net/http"
	"testing"
	
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type errorHTTPSuite struct {
	suite.Suite
}

func TestErrorHTTPInit(t *testing.T) {
	suite.Run(t, new(errorHTTPSuite))
}

func (r *errorHTTPSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *errorHTTPSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *errorHTTPSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *errorHTTPSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *errorHTTPSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *errorHTTPSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *errorHTTPSuite) TestErrorHTTP_ToJSON() {
	Convey("Given a error HTTP ", r.T(), func() {
		err := errorAVA.ErrorHTTP{
			HTTPStatus: http.StatusOK,
			Code:       -1,
			Message:    http.StatusText(http.StatusOK),
		}
		So(err.ToJSON(), ShouldResemble, "{\"status\":200,\"code\":-1,\"message\":\"OK\"}")
	})
}

func (r *errorHTTPSuite) TestErrorHTTP_Error() {
	Convey("Given a error HTTP ", r.T(), func() {
	
	})
}

func (r *errorHTTPSuite) TestErrorHTTP_WriteToResponse() {
	Convey("Given a error HTTP ", r.T(), func() {
	
	})
}
