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
	// ApacheErrorFormatTemplate : [{timestamp}] [{module}:{severity}] [pid {pid}:tid {threadID}] [client: %{client}] %{message}
	ApacheErrorFormatTemplate = fmt.Sprintf("[{%s}] [{%s}:{%s}] [pid {%s}:tid {%s}] [client: %s[%s]s] %s[%s]s\n",
		TIMESTAMP, MODULE, SEVERITY, PROCESS_PID, THREAD_ID, "%", CLIENT, "%", MESSAGE)

	// message template just logs the message.
	ApacheErrorMessageTemplate = fmt.Sprintf("%s[%s]s\n", "%", MESSAGE)

	// Detailed template logs padded columns including the running PID.
	ApacheErrorDetailedTemplate = "%[ASC_TIME]s %-5[process]d %-7[LEVEL_NAME]s %-20[NAME]s %[MESSAGE]s%[FIELDS]s\n"

	// defaultTimestampFormat is the default format used if the user does not specify their own.
	ApacheErrorDefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type ApacheErrorLogFormatter struct {
	customFormatter
}

func init() {
	loggerAVA.RegisterLogger(loggerAVA.LogFormatterTypeApacheErrorLog, NewLoggerApacheErrorLogFormatterDefault())
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func newApacheErrorLogFormatter() *ApacheErrorLogFormatter {
	var disableColors bool = false
	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		disableColors = true
	}
	
	formatter := ApacheErrorLogFormatter{struct {
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
		template:        ApacheErrorFormatTemplate,
		handlers:        nil,
		attributes:      nil,
		forceColors:     false,
		disableColors:   disableColors,
		timestampFormat: ApacheErrorDefaultTimestampFormat,
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
	formatter.parseTemplate(ApacheErrorFormatTemplate, nil)
	
	logrus.SetFormatter(&formatter)
	
	return &formatter
}

func NewLoggerApacheErrorLogFormatterDefault() *loggerAVA.Logger {
	logger := loggerAVA.Logger{
		Level:     loggerAVA.LogLevelTypeDebug,
		Formatter: loggerAVA.LogFormatterTypeApacheErrorLog,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	logger.Log.SetFormatter(newApacheErrorLogFormatter())
	logger.Log.SetOutput(os.Stdout)
	logger.Log.SetLevel(logrus.DebugLevel)
	
	return &logger
}
