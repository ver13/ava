package formatter

import (
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

const (
	// RFC3164LogTemplate  : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3164LogTemplate = "<%-7[LEVEL_NAME]s>%[ASC_TIME]s %[hostname]s %[application]s[%-5[process]d]: %[MESSAGE]s"

	// message template just logs the message.
	RFC3164MessageTemplate = "%[MESSAGE]s\n"

	// Detailed template logs padded columns including the running PID.
	RFC3164DetailedTemplate = "%[ASC_TIME]s %-5[process]d %-7[LEVEL_NAME]s %-20[NAME]s %[MESSAGE]s%[FIELDS]s\n"

	// defaultTimestampFormat is the default format used if the user does not specify their own.
	RFC3164DefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type RFC3164Formatter struct {
	customFormatter
}

func init() {
	loggerAVA.RegisterLogger(loggerAVA.LogFormatterTypeRFC3164Log, NewLoggerRFC3164Default())
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func newRFC3164Formatter() *RFC3164Formatter {
	formatter := RFC3164Formatter{
		customFormatter: customFormatter{
			colorDebug:      AnsiCyan,
			colorInfo:       AnsiGreen,
			colorWarn:       AnsiYellow,
			colorError:      AnsiRed,
			colorFatal:      AnsiMagenta,
			colorPanic:      AnsiMagenta,
			timestampFormat: RFC3164DefaultTimestampFormat,
			startTime:       time.Now(),
		},
	}

	// Parser the template string.
	formatter.parseTemplate(RFC3164LogTemplate, nil)

	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		formatter.disableColors = true
	}

	return &formatter
}

func NewLoggerRFC3164Default() *loggerAVA.Logger {
	logger := loggerAVA.Logger{
		Level:     loggerAVA.LogLevelTypeDebug,
		Formatter: loggerAVA.LogFormatterTypeRFC3164Log,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	logger.Log.SetFormatter(newRFC3164Formatter())
	logger.Log.SetOutput(os.Stdout)
	logger.Log.SetLevel(logrus.DebugLevel)

	return &logger
}
