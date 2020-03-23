package logger

import (
	"fmt"
	"sync"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorLoggerAVA "github.com/ver13/ava/pkg/common/logger/error"
)

type loggerFactory map[LogFormatterType]*Logger

var (
	factory loggerFactory
	once    sync.Once
	active  LogFormatterType
)

func init() {
	once.Do(func() {
		factory = make(map[LogFormatterType]*Logger)
		active = LogFormatterTypeUnknown
	})
}

func RegisterLogger(t LogFormatterType, l *Logger) *errorAVA.Error {
	factory[t] = l
	return nil
}

func UseLogger(t LogFormatterType) (*Logger, *errorAVA.Error) {
	logger := factory[t]
	if logger == nil {
		return nil, errorLoggerAVA.NotImplemented(nil, fmt.Sprintf("Logger type: %s", t.String()))
	}
	active = t
	return logger, nil
}

func GetInstance() *Logger {
	return factory[active]
}
