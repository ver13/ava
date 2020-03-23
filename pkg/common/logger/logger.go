package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorLoggerAVA "github.com/ver13/ava/pkg/common/logger/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

type Logger struct {
	Enable bool

	Level     LogLevelType
	Formatter LogFormatterType

	Output struct {
		Console bool
		File    bool
	}

	File *lumberjack.Logger

	Log *logrus.Logger

	Mux sync.Mutex
}

func (l *Logger) IsValid() *errorAVA.Error {
	if l.File != nil {
		if l.File.Filename == "" {
			return errorLoggerAVA.InvalidConfig(nil, "Filename is empty.")
		}

		// maxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename.
		if l.File.MaxAge > 30 {
			l.File.MaxAge = 30
		}

		// MaxSize is the maximum size in megabytes of the log File before it gets rotated.
		if l.File.MaxSize > 10 {
			l.File.MaxSize = 10
		}

		// MaxBackups is the maximum number of old log files to retain.
		if l.File.MaxBackups > 5 {
			l.File.MaxBackups = 5
		}
	}

	return nil
}

func (l *Logger) Reset() {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	l = &Logger{
		Level:     LogLevelTypeDebug,
		Formatter: LogFormatterTypeText,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	l.Log.SetFormatter(&logrus.TextFormatter{})
	l.Log.SetOutput(os.Stdout)
	l.Log.SetLevel(logrus.DebugLevel)
}

// debug logs a message at level debug on the standard logger.
func (l *Logger) Debug(args ...interface{}) {
	l.Log.Debug(args...)
}
func Debug(args ...interface{}) {
	GetInstance().Debug(args...)
}

// debug logs a message at level debug on the standard logger.
func (l *Logger) Debugln(args ...interface{}) {
	l.Log.Debugln(args...)
}
func Debugln(args ...interface{}) {
	GetInstance().Debugln(args...)
}

// Debugf logs a message at level debug on the standard logger.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Log.Debugf(format, args...)
}
func Debugf(format string, args ...interface{}) {
	GetInstance().Debugf(format, args...)
}

// info logs a message at level info on the standard logger.
func (l *Logger) Info(args ...interface{}) {
	l.Log.Info(args...)
}
func Info(args ...interface{}) {
	GetInstance().Info(args...)
}

// info logs a message at level info on the standard logger.
func (l *Logger) Infoln(args ...interface{}) {
	l.Log.Infoln(args...)
}
func Infoln(args ...interface{}) {
	GetInstance().Infoln(args...)
}

// Infof logs a message at level info on the standard logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Log.Infof(format, args...)
}
func Infof(format string, args ...interface{}) {
	GetInstance().Infof(format, args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.Log.Trace(args...)
}
func Trace(args ...interface{}) {
	GetInstance().Trace(args...)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.Log.Traceln(args...)
}
func Traceln(args ...interface{}) {
	GetInstance().Traceln(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Log.Tracef(format, args...)
}
func Tracef(format string, args ...interface{}) {
	GetInstance().Tracef(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l *Logger) Warn(args ...interface{}) {
	l.Log.Warn(args...)
}
func Warn(args ...interface{}) {
	GetInstance().Warn(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l *Logger) Warnln(args ...interface{}) {
	l.Log.Warnln(args...)
}
func Warnln(args ...interface{}) {
	GetInstance().Warnln(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Log.Warnf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	GetInstance().Warnf(format, args...)
}

// err logs a message at level err on the standard logger.
func (l *Logger) Error(args ...interface{}) {
	l.Log.Error(args...)
}
func Error(args ...interface{}) {
	GetInstance().Error(args...)
}

// err logs a message at level err on the standard logger.
func (l *Logger) Errorln(args ...interface{}) {
	l.Log.Errorln(args...)
}
func Errorln(args ...interface{}) {
	GetInstance().Errorln(args...)
}

// Errorf logs a message at level err on the standard logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Log.Errorf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	GetInstance().Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l *Logger) Fatal(args ...interface{}) {
	l.Log.Fatal(args...)
}
func Fatal(args ...interface{}) {
	GetInstance().Fatal(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l *Logger) Fatalln(args ...interface{}) {
	l.Log.Fatalln(args...)
}
func Fatalln(args ...interface{}) {
	GetInstance().Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Log.Fatalf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	GetInstance().Fatalf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l *Logger) Panic(args ...interface{}) {
	l.Log.Panic(args...)
}
func Panic(args ...interface{}) {
	GetInstance().Panic(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l *Logger) Panicln(args ...interface{}) {
	l.Log.Panicln(args...)
}
func Panicln(args ...interface{}) {
	GetInstance().Panicln(args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Log.Panicf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	GetInstance().Panicf(format, args...)
}

func (l *Logger) With(key string, value interface{}) LoggerI {
	return &Logger{
		Log: l.Log.WithField(key, value).Logger,
	}
}
func With(key string, value interface{}) {
	GetInstance().With(key, value)
}

// WithFields takes into consideration the fields
func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Log.WithFields(fields)
}
func WithFields(fields logrus.Fields) *logrus.Entry {
	return GetInstance().WithFields(fields)
}

func (l *Logger) IsEnable() bool {
	return l.Enable
}
func IsEnable() bool {
	return GetInstance().IsEnable()
}

func (l *Logger) SetEnable() {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	l.Enable = true
}
func SetEnable() {
	GetInstance().SetEnable()
}

func (l *Logger) GetFormat() LogFormatterType {
	return l.Formatter
}
func GetFormat() LogFormatterType {
	return GetInstance().GetFormat()
}

func (l *Logger) SetFormat(logFormatterType LogFormatterType) *errorAVA.Error {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	formatter, err := UseLogger(logFormatterType)
	if err != nil {
		return err
	}
	l = formatter

	return nil
}
func SetFormat(logFormatterType LogFormatterType) *errorAVA.Error {
	return GetInstance().SetFormat(logFormatterType)
}

func (l *Logger) GetLevel() LogLevelType {
	return l.Level
}
func GetLevel() LogLevelType {
	return GetInstance().GetLevel()
}

func (l *Logger) SetLevel(levelType LogLevelType) *errorAVA.Error {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	lvl, err := logrus.ParseLevel(levelType.String())
	if err != nil {
		return errorLoggerAVA.LevelWrong(err, levelType.String())
	}

	l.Log.Level = lvl
	return nil
}
func SetLevel(levelType LogLevelType) *errorAVA.Error {
	return GetInstance().SetLevel(levelType)
}

// SetOutput sets the logger output.
func (l *Logger) SetOutputConsole(output io.Writer) {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	l.Output.Console = true

	if l.Output.File {
		l.Log.SetOutput(io.MultiWriter(output, l.File))
	} else {
		l.Log.SetOutput(output)
	}
}
func SetOutputConsole(output io.Writer) {
	GetInstance().SetOutputConsole(output)
}

// SetOutput sets the logger output.
func (l *Logger) SetOutputFile(filename string,
	maxSize int,
	maxBackups int,
	maxAge int,
	compress bool) {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	l.Output.File = true
	l.File = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  false,
		Compress:   compress,
	}

	if l.Output.Console {
		l.Log.SetOutput(io.MultiWriter(os.Stdout, l.File))
	} else {
		l.Log.SetOutput(io.MultiWriter(l.File))
	}
}
func SetOutputFile(filename string,
	maxSize int,
	maxBackups int,
	maxAge int,
	compress bool) {
	GetInstance().SetOutputFile(filename, maxSize, maxBackups, maxAge, compress)
}

func (l *Logger) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	if s := serializerAVA.GetSerializer(t); s != nil {
		return s.Serializer(l)
	}
	return nil, errorSerializerAVA.NotImplemented(nil, fmt.Sprintf("Serializer type: %s", t.String()))
}

func (l *Logger) Parser() (*Logger, *errorAVA.Error) {
	panic("Not implemented.")
}

func Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	return GetInstance().Serializer(t)
}

