package formatter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

// ANSI color codes.
const (
	AnsiReset     = 0
	AnsiRed       = 31
	AnsiHiRed     = 91
	AnsiGreen     = 32
	AnsiHiGreen   = 92
	AnsiYellow    = 33
	AnsiHiYellow  = 93
	AnsiBlue      = 34
	AnsiHiBlue    = 94
	AnsiMagenta   = 35
	AnsiHiMagenta = 95
	AnsiCyan      = 36
	AnsiHiCyan    = 96
	AnsiWhite     = 37
	AnsiHiWhite   = 97
)

var _reAnsi = regexp.MustCompile("\033\\[[\\d;]+m")

// From fmt.Sprintf()'s source code. Only handle one format (%2s) for one value. `a` width handled by caller.
func sprintfColorString(format, a string, w int) string {
	pos := 1 // format should always start with %.
	minus := false
	buffer := []byte{}
	defer func() { buffer = buffer[:0] }()

	// Handle '-' character after %.
	if format[pos] == '-' {
		minus = true
		pos++
	}

	// Parser padding/width.
	width := 0
	for '0' <= format[pos] && format[pos] <= '9' {
		width = width*10 + int(format[pos]-'0')
		pos++
	}
	width -= w
	if width <= 0 {
		buffer = append(buffer, a...)
		return string(buffer)
	}

	// Add padding.
	if minus {
		buffer = append(buffer, a...)
	}
	padding := make([]byte, width)
	for i := 0; i < width; i++ {
		padding[i] = byte(' ')
	}
	buffer = append(buffer, padding...)
	if !minus {
		buffer = append(buffer, a...)
	}
	return string(buffer)
}

// Sprintf is like fmt.Sprintf() but exclude ANSI color sequences from string padding.
func (f *customFormatter) Sprintf(values ...interface{}) string {
	if (!f.forceColors && f.disableColors) || len(f.handleColors) == 0 {
		return fmt.Sprintf(f.template, values...)
	}
	template := f.template
	for i := len(f.handleColors) - 1; i >= 0; i-- {
		value := values[f.handleColors[i][0]].(string)
		pos := strings.Index(value, "\033")
		if pos < 0 {
			continue
		}

		// Determine width without color sequences.
		width := utf8.RuneCountInString(value)
		for _, p := range _reAnsi.FindAllStringIndex(value[pos:], -1) {
			width -= p[1] - p[0]
		}

		// Pull formatting from template.
		start, end := f.handleColors[i][1], f.handleColors[i][2]
		format := template[start:end]
		template = template[:start] + "%s" + template[end:]

		// format value while not counting ANSI color codes (yet still including them).
		values[f.handleColors[i][0]] = sprintfColorString(format, value, width)
	}
	return fmt.Sprintf(template, values...)
}

// Color colorizes the input string and returns it with ANSI color codes.
func Color(entry *logrus.Entry, formatter CustomFormatterI, s string) string {
	if !formatter.IsForceColors() && formatter.IsDisableColors() {
		return s
	}

	// Determine color. Default is info.
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = formatter.Color(loggerAVA.LogLevelTypeDebug)
	case logrus.WarnLevel:
		levelColor = formatter.Color(loggerAVA.LogLevelTypeWarn)
	case logrus.ErrorLevel:
		levelColor = formatter.Color(loggerAVA.LogLevelTypeError)
	case logrus.PanicLevel:
		levelColor = formatter.Color(loggerAVA.LogLevelTypePanic)
	case logrus.FatalLevel:
		levelColor = formatter.Color(loggerAVA.LogLevelTypeFatal)
	default:
		levelColor = formatter.Color(loggerAVA.LogLevelTypeInfo)
	}
	if levelColor == AnsiReset {
		return s
	}

	// Colorize.
	return "\033[" + strconv.Itoa(levelColor) + "m" + s + "\033[0m"
}

// WindowsNativeANSI returns true if either the stderr or stdout consoles natively support ANSI color codes. On
// non-Windows platforms this always returns false.
func WindowsNativeANSI() bool {
	enabled, _ := windowsNativeANSI(true, false, nil)
	if enabled {
		return enabled
	}
	enabled, _ = windowsNativeANSI(false, false, nil)
	return enabled
}

// WindowsEnableNativeANSI will attempt to set ENABLE_VIRTUAL_TERMINAL_PROCESSING on a console using SetConsoleMode.
//
// :param stdErr: Issue SetConsoleMode win32 api call on stderr instead of stdout handle.
func WindowsEnableNativeANSI(stderr bool) *errorAVA.Error {
	_, err := windowsNativeANSI(stderr, true, nil)
	return err
}
