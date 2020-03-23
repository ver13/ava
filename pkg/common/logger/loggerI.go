package logger

import (
	"io"

	"github.com/sirupsen/logrus"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

// logger is the interface for loggers used in the AVA components.
type LoggerI interface {
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})

	Trace(...interface{})
	Traceln(...interface{})
	Tracef(string, ...interface{})

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})

	Panic(...interface{})
	Panicln(...interface{})
	Panicf(format string, args ...interface{})

	WithFields(logrus.Fields) *logrus.Entry

	With(key string, value interface{}) LoggerI

	SetOutputConsole(output io.Writer)
	SetOutputFile(filename string, maxSize int, maxBackups int, maxAge int, compress bool)

	SetFormat(logFormatterType LogFormatterType) *errorAVA.Error
	GetFormat() LogFormatterType

	IsEnable() bool
	SetEnable()

	SetLevel(levelType LogLevelType) *errorAVA.Error
	GetLevel() LogLevelType

	Reset()

	Serializer(serializerAVA.SerializerType) ([]byte, *errorAVA.Error)
}
