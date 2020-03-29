package file

import (
	"fmt"
	"io"
	"os"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type LoggerConfig struct {
	format string `mapstructure:"format,omitempty"`
	level  string `mapstructure:"level,omitempty"`
	enable bool   `mapstructure:"enable,omitempty"`
	output struct {
		console bool `mapstructure:"console,omitempty"`
		file    bool `mapstructure:"file,omitempty"`
	} `mapstructure:"output,omitempty"`
	file struct {
		// Filename is the file to write logs to.  Backup logger files will be retained in the same directory.  It uses <processname>-lumberjack.logger in os.TempDir() if empty.
		filename string `mapstructure:"filename"`

		// MaxSize is the maximum size in megabytes of the logger file before it gets rotated. It defaults to 100 megabytes.
		maxSize int `mapstructure:"maxsize"`

		// maxAge is the maximum number of days to retain old logger files based on the timestamp encoded in their filename.  Note that a day is defined as 24 hours and may not exactly correspond to calendar days due to daylight savings, leap seconds, etc. The default is not to remove old logger files based on age.
		maxAge int `mapstructure:"maxage"`

		// MaxBackups is the maximum number of old logger files to retain.  The default is to retain all old logger files (though maxAge may still cause them to get deleted.)
		maxBackups int `mapstructure:"maxbackups"`

		// LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time.  The default is to use UTC time.
		localTime bool `mapstructure:"localtime"`

		// Compress determines if the rotated logger files should be compressed using gzip. The default is not to perform compression.
		compress bool `mapstructure:"compress"`
	} `mapstructure:"file,omitempty"`
}

func (l *LoggerConfig) Parser() (*loggerAVA.Logger, *errorAVA.Error) {
	logLevel, err := loggerAVA.ParseLogLevelType(l.level)
	if err != nil {
		fmt.Errorf("loggerAVA level incorrect in configurationServiceI file." + l.level)
		fmt.Errorf("loggerAVA level = [DEBUG]")
		logLevel = loggerAVA.LogLevelTypeDebug
	}

	logType, err := loggerAVA.ParseLogFormatterType(l.format)
	if err != nil {
		fmt.Errorf("loggerAVA type incorrect in configurationServiceI file." + l.format)
		fmt.Errorf("loggerAVA format = " + logType.String())
		logType = loggerAVA.LogFormatterTypeText
	}

	logger := loggerAVA.GetInstance()
	logger.SetFormat(logType)
	if l.enable {
		logger.SetEnable()
	}
	logger.SetLevel(logLevel)
	if l.output.console {
		logger.SetOutputConsole(io.MultiWriter(os.Stdout, os.Stderr))
	}
	logger.SetOutputFile(l.file.filename, l.file.maxSize, l.file.maxBackups, l.file.maxAge, l.file.compress)

	return logger, nil
}

func (l *LoggerConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	s, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return s.Serializer(l)
}
