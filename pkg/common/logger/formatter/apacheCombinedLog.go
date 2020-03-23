package formatter

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

var (
	// Basic template just logs the level name, name field, message and fields.
	// ApacheCombinedFormatTemplate = "%[host]s - %[user]s %[authUserId]d [%[ASC_TIME]s] \"%[method]s %[request]s HTTP/1.0\" %[responseCode]d %[bytes]d \"%[referrer]s\" \"%[agent]s\""
	ApacheCombinedFormatTemplate = fmt.Sprintf("%s[%s]s - %s[%s]s %s[%s]d [%s[%s]s] \"%s[%s]s %s[%s]s HTTP/1.0\" %s[%s]d %s[%s]d \"%s[%s]s\" \"%s[%s]s\"\n",
	"%", HOST, "%", USER, "%", AUTH_USER_ID, "%", ASC_TIME, "%", METHOD, "%", REQUEST, "%", RESPONSE_CODE, "%", BYTES,"%",REF_ERRER,"%",AGENT)

	// message template just logs the message.
	ApacheCombinedMessageTemplate = fmt.Sprintf("%s[%s]s\n", "%", MESSAGE)

	// Detailed template logs padded columns including the running PID.
	ApacheCombinedDetailedTemplate = "%[ASC_TIME]s %-5[process]d %-7[LEVEL_NAME]s %-20[NAME]s %[message]s%[FIELDS]s\n"

	// defaultTimestampFormat is the default format used if the user does not specify their own.
	ApacheCombinedDefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type ApacheCombinedFormatter struct {
	customFormatter
}

func init() {
	loggerAVA.RegisterLogger(loggerAVA.LogFormatterTypeApacheCombinedLog, NewLoggerApacheCombinedLogFormatterDefault())
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func newApacheCombinedLogFormatter() *ApacheCombinedFormatter {
	var disableColors bool = false
	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		disableColors = true
	}

	formatter := ApacheCombinedFormatter{struct {
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
		template:        ApacheCombinedFormatTemplate,
		handlers:        nil,
		attributes:      nil,
		forceColors:     false,
		disableColors:   disableColors,
		timestampFormat: ApacheCombinedDefaultTimestampFormat,
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
	formatter.parseTemplate(ApacheCombinedFormatTemplate, nil)

	return &formatter
}

func NewLoggerApacheCombinedLogFormatterDefault() *loggerAVA.Logger {
	logger := loggerAVA.Logger{
		Level:     loggerAVA.LogLevelTypeDebug,
		Formatter: loggerAVA.LogFormatterTypeApacheCombinedLog,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	logger.Log.SetFormatter(newApacheCombinedLogFormatter())
	logger.Log.SetOutput(os.Stdout)
	logger.Log.SetLevel(logrus.DebugLevel)
	
	return &logger
}
