package formatter

import (
	"os"
	"runtime"
	"time"
	
	"github.com/sirupsen/logrus"
	
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

const (
	// RFC5424LogTemplate  : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	RFC5424LogTemplate = "<priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}"
	
	// message template just logs the message.
	RFC5424MessageTemplate = "%[MESSAGE]s\n"
	
	// Detailed template logs padded columns including the running PID.
	RFC5424DetailedTemplate = "%[ASC_TIME]s %-5[process]d %-7[LEVEL_NAME]s %-20[NAME]s %[MESSAGE]s%[FIELDS]s\n"
	
	// defaultTimestampFormat is the default format used if the user does not specify their own.
	RFC5424DefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type RFC5424Formatter struct {
	customFormatter
}

func init() {
	loggerAVA.RegisterLogger(loggerAVA.LogFormatterTypeRFC5424Log, NewLoggerRFC5424Default())
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func newRFC5424Formatter() *RFC5424Formatter {
	formatter := RFC5424Formatter{
		customFormatter: customFormatter{
			colorDebug:      AnsiCyan,
			colorInfo:       AnsiGreen,
			colorWarn:       AnsiYellow,
			colorError:      AnsiRed,
			colorFatal:      AnsiMagenta,
			colorPanic:      AnsiMagenta,
			timestampFormat: RFC5424DefaultTimestampFormat,
			startTime:       time.Now(),
		},
	}
	
	// Parser the template string.
	formatter.parseTemplate(RFC5424LogTemplate, nil)
	
	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		formatter.disableColors = true
	}
	
	return &formatter
}

func NewLoggerRFC5424Default() *loggerAVA.Logger {
	logger := loggerAVA.Logger{
		Level:     loggerAVA.LogLevelTypeDebug,
		Formatter: loggerAVA.LogFormatterTypeRFC5424Log,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	logger.Log.SetFormatter(newRFC5424Formatter())
	logger.Log.SetOutput(os.Stdout)
	logger.Log.SetLevel(logrus.DebugLevel)
	
	return &logger
}
