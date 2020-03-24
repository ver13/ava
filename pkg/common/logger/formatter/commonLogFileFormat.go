package formatter

import (
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

const (
	// CommonLogFileFormatTemplate : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} HTTP/1.0" {response-code} {bytes}
	CommonFileFormatTemplate = "%[host]s - %[user]s %[authUserId]d [%[ASC_TIME]s] \"%[method]s %[request]s HTTP/1.0\" %[responseCode]d %[bytes]d"

	// message template just logs the message.
	CommonFileFormatMessageTemplate = "%[MESSAGE]s\n"

	// Detailed template logs padded columns including the running PID.
	CommonFileFormatDetailedTemplate = "%[ASC_TIME]s %-5[process]d %-7[LEVEL_NAME]s %-20[NAME]s %[MESSAGE]s%[FIELDS]s\n"

	// defaultTimestampFormat is the default format used if the user does not specify their own.
	CommonFileFormatDefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type CommonLogFileFormatter struct {
	customFormatter
}

func init() {
	loggerAVA.RegisterLogger(loggerAVA.LogFormatterTypeCommonLogFileFormat, NewLoggerCommonLogfileFormatterDefault())
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func newCommonLogfileFormatter() *CommonLogFileFormatter {
	var disableColors = false
	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		disableColors = true
	}

	formatter := CommonLogFileFormatter{struct {
		template        string
		handlers        []handler
		attributes      attributes
		forceColors     bool
		disableColors   bool
		timestampFormat string
		disableSorting  bool
		colorDebug      int
		colorInfo       int
		colorWarn       int
		colorError      int
		colorFatal      int
		colorPanic      int
		color           map[loggerAVA.LogLevelType]int
		handleColors    [][3]int
		startTime       time.Time
	}{
		template:        CommonFileFormatTemplate,
		handlers:        nil,
		attributes:      nil,
		forceColors:     false,
		disableColors:   disableColors,
		timestampFormat: CommonFileFormatDefaultTimestampFormat,
		disableSorting:  true,
		colorDebug:      AnsiCyan,
		colorInfo:       AnsiGreen,
		colorWarn:       AnsiYellow,
		colorError:      AnsiRed,
		colorFatal:      AnsiMagenta,
		colorPanic:      AnsiMagenta,
		color:           nil,
		handleColors:    nil,
		startTime:       time.Now(),
	},
	}

	// Parser the template string.
	formatter.parseTemplate(CommonFileFormatTemplate, nil)

	return &formatter
}

func NewLoggerCommonLogfileFormatterDefault() *loggerAVA.Logger {
	logger := loggerAVA.Logger{
		Level:     loggerAVA.LogLevelTypeDebug,
		Formatter: loggerAVA.LogFormatterTypeCommonLogFileFormat,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	logger.Log.SetFormatter(newCommonLogfileFormatter())
	logger.Log.SetOutput(os.Stdout)
	logger.Log.SetLevel(logrus.DebugLevel)

	return &logger
}
