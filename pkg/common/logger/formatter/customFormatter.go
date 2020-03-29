package formatter

import (
	"bytes"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/ver13/ava/pkg/common/logger"
)

const (
	// defaultTimestampFormat is the default format used if the user does not specify their own.
	defaultTimestampFormat = "2006-01-02 15:04:05.000"
)

// customFormatter is the main formatter for the library.
type customFormatter struct {
	// Post-processed formatting template (e.g. "%s:%s:%s\n").
	template string

	// handler functions whose indexes match up with template Sprintf explicit argument indexes.
	handlers []handler

	// Attribute names (e.g. "levelName") used in pre-processed template.
	attributes attributes

	// Set to true to bypass checking for a TTY before outputting colors.
	forceColors bool

	// Force disabling colors and bypass checking for a TTY.
	disableColors bool

	// Timestamp format %[ascTime]s will use for display when a full timestamp is printed.
	timestampFormat string

	// The fields are sorted by default for a consistent output. For applications that log extremely frequently this may not be desired.
	disableSorting bool

	// Different colors for different log levels.
	colorDebug int
	colorInfo  int
	colorWarn  int
	colorError int
	colorFatal int
	colorPanic int

	color map[logger.LogLevelType]int

	handleColors [][3]int
	startTime    time.Time
}

func (f *customFormatter) SetColor(t logger.LogLevelType, reset int) {
	f.color[t] = reset
}

func (f *customFormatter) EnableForceColors() {
	f.forceColors = true
}

func (f *customFormatter) IsForceColors() bool {
	return f.forceColors
}

func (f *customFormatter) IsDisableColors() bool {
	return f.disableColors
}

func (f *customFormatter) Color(t logger.LogLevelType) int {
	return f.color[t]
}

// format is called by logrus and returns the formatted string.
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Call handlers.
	values := make([]interface{}, len(f.handlers))
	for i, handler := range f.handlers {
		value, err := handler(entry, f)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	// Parser template and return.
	parsed := f.Sprintf(values...)
	return bytes.NewBufferString(parsed).Bytes(), nil
}

// NewFormatter creates a new customFormatter, sets the template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func NewFormatter(template string, custom CustomHandlers) CustomFormatterI {
	formatter := customFormatter{
		colorDebug:      AnsiCyan,
		colorInfo:       AnsiGreen,
		colorWarn:       AnsiYellow,
		colorError:      AnsiRed,
		colorFatal:      AnsiMagenta,
		colorPanic:      AnsiMagenta,
		timestampFormat: defaultTimestampFormat,
		startTime:       time.Now(),
	}

	// Parser the template string.
	formatter.parseTemplate(template, custom)

	// Disable colors if not supported.
	if !checkIfTerminal(logrus.StandardLogger().Out) || (runtime.GOOS == "windows" && !WindowsNativeANSI()) {
		formatter.disableColors = true
	}

	return &formatter
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
