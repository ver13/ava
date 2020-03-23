package logger_test

import (
	"os"
	"testing"
	
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/logger"
)

type loggerAVASuite struct {
	suite.Suite
}

func TestLoggerAVAInit(t *testing.T) {
	suite.Run(t, new(loggerAVASuite))
}

func (log *loggerAVASuite) BeforeTest() {
	log.T().Log("BeforeTest")
}

func (log *loggerAVASuite) AfterTest() {
	log.T().Log("AfterTest")
}

func (log *loggerAVASuite) SetupSuite() {
	log.T().Log("SetupSuite")
}

func (log *loggerAVASuite) SetupTest() {
	log.T().Log("SetupTest")
}

func (log *loggerAVASuite) TearDownSuite() {
	log.T().Log("TearDownSuite")
}

func (log *loggerAVASuite) TearDownTest() {
	log.T().Log("TearDownTest")
}

func (log *loggerAVASuite) TestLoggerGmf_New() {
	Convey("Given create a logger", log.T(), func() {
		Convey("When logger configurationService is ok", func() {
			stdout, stderr, err := WithCapSys(func() {
				logger := GetInstance()
				logger.SetOutputConsole(os.Stdout)
				logger.Debug("debug type trace.")
			})

			So(err, ShouldBeNil)
			So(stdout, ShouldNotBeEmpty)
			So(stderr, ShouldBeEmpty)
		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Debug() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Debugln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Debugf() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Trace() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Traceln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Tracef() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Info() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Infoln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Infof() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Warn() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Warnln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Warnf() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Error() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Errorln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Errorf() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Fatal() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Fatalln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Fatalf() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Panic() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Panicln() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (log *loggerAVASuite) TestLoggerGmf_Panicf() {
	Convey("Given a ", log.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
