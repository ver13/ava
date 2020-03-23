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
	// Text log template just logs the level name, name field, message and fields.
	//	TextLogTemplate = "%[levelName]s:%[name]s:%[message]s%[fields]s\n"
	TextLogTemplate = fmt.Sprintf("%s[%s]s:%s[%s]s:%s[%s]s:%s[%s]s\n", "%", LEVEL_NAME, "%", NAME, "%", MESSAGE, "%", FIELDS)

	// message template just logs the message.
	TextMessageTemplate = fmt.Sprintf("%s[%s]s\n", "%", PROCESS_PID)

	// Detailed template logs padded columns including the running PID.
	TestDetailedTemplate = "%[ASC_TIME]s %-5[process]d %-7[LEVEL_NAME]s %-20[NAME]s %[MESSAGE]s%[FIELDS]s\n"

	// defaultTimestampFormat is the default format used if the user does not specify their own.
	TextDefaultTimestampFormat = "2006-01-02 15:04:05.000"
)

type TextFormatter struct {
	customFormatter
}

func init() {
	loggerAVA.RegisterLogger(loggerAVA.LogFormatterTypeText, NewLoggerTextFormatterDefault())
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func newTextFormatter() *TextFormatter {
	var disableColors bool = false
	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		disableColors = true
	}
	
	formatter := TextFormatter{struct {
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
			template: TextLogTemplate,
			handlers: nil,
			attributes: nil,
			forceColors: false,
			disableColors: disableColors,
			timestampFormat: TextDefaultTimestampFormat,
			disableSorting: true,
			colorDebug: AnsiCyan,
			colorInfo: AnsiGreen,
			colorWarn: AnsiYellow,
			colorError: AnsiRed,
			colorFatal: AnsiMagenta,
			colorPanic: AnsiMagenta,
			color: nil,
			handleColors: nil,
			startTime: time.Now(),
		},
	}
	
	// Parser the template string.
	formatter.parseTemplate(TextLogTemplate, nil)
	
	return &formatter
}

func NewLoggerTextFormatterDefault() *loggerAVA.Logger {
	logger := loggerAVA.Logger{
		Level:     loggerAVA.LogLevelTypeDebug,
		Formatter: loggerAVA.LogFormatterTypeText,
		Enable:    true,
		Output: struct {
			Console bool
			File    bool
		}{true, false},
		File: nil,
		Log:  logrus.New(),
	}
	logger.Log.SetFormatter(newTextFormatter())
	logger.Log.SetOutput(os.Stdout)
	logger.Log.SetLevel(logrus.DebugLevel)
	
	return &logger
}
