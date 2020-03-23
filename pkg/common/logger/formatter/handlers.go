package formatter

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var _reBracketed = regexp.MustCompile(`%([\d.-]*)\[(\w+)](\w)`)

// handler is the function signature of formatting attributes such as "levelName" and "message".
type handler func(*logrus.Entry, *customFormatter) (interface{}, error)

// CustomHandlers is a mapping of handler-type functions to attributes as key names (e.g. "levelName").
//
// With this type many custom handler functions can be defined and fed to NewFormatter(). CustomHandlers are parsed
// first so you can override built-in handlers such as the one for %[ASC_TIME]s with your own. Since they are exported
// you can call built-in handlers in your own custom handler. The returned interface{} value is passed to fmt.Sprintf().
//
// In addition to overriding handlers you can create new attributes (such as %[myAttr]s) and map it to your handler
// function.
type CustomHandlers map[string]handler

// attributes is a map used like a "set" to keep track of which formatting attributes are used.
type attributes map[string]bool

// Contains returns true if attr is present.
func (a attributes) Contains(attr string) bool {
	_, ok := a[attr]
	return ok
}

// handlerAscTime returns the formatted timestamp of the entry.
func handlerAscTime(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return entry.Time.Format(formatter.timestampFormat), nil
}

// handlerFields returns the entry's fields (excluding name field if %[name]s is used) colorized according to log level.
// Fields' formatting: key=value key2=value2
func handlerFields(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	var fields string

	// Without sorting no need to get keys from map into a string array.
	if formatter.disableSorting {
		for key, value := range entry.Data {
			if key == "name" && formatter.attributes.Contains("name") {
				continue
			}
			fields = fmt.Sprintf("%s %s=%v", fields, Color(entry, formatter, key), value)
		}
		return fields, nil
	}

	// Put keys in a string array and sort it.
	keys := make([]string, len(entry.Data))
	i := 0
	for k := range entry.Data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// Do the rest.
	for _, key := range keys {
		if key == "name" && formatter.attributes.Contains("name") {
			continue
		}
		fields = fmt.Sprintf("%s %s=%v", fields, Color(entry, formatter, key), entry.Data[key])
	}
	return fields, nil
}

// handlerLevelName returns the entry's long level name (e.g. "WARNING").
func handlerLevelName(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, strings.ToUpper(entry.Level.String())), nil
}

// handlerName returns the name field value set by the user in entry.Data.
func handlerName(entry *logrus.Entry, _ *customFormatter) (interface{}, error) {
	if value, ok := entry.Data["name"]; ok {
		return value.(string), nil
	}
	return "", nil
}

// handlerMessage returns the unformatted log message in the entry.
func handlerMessage(entry *logrus.Entry, _ *customFormatter) (interface{}, error) {
	return entry.Message, nil
}

// handlerProcessPID returns the current process' PID.
func handlerProcessPID(_ *logrus.Entry, _ *customFormatter) (interface{}, error) {
	return os.Getpid(), nil
}

// handlerRelativeCreated returns the number of seconds since program start time.
func handlerRelativeCreated(_ *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return int(time.Since(formatter.startTime) / time.Second), nil
}

// handlerShortLevelName returns the first 4 letters of the entry's level name (e.g. "WARN").
func handlerShortLevelName(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, strings.ToUpper(entry.Level.String()[:4])), nil
}

// parseTemplate parses the template string and prepares it for fmt.Sprintf() and keeps track of which handlers to use.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
func (f *customFormatter) parseTemplate(template string, custom CustomHandlers) {
	f.attributes = make(attributes)
	segments := []string{}
	segmentsPos := 0

	for _, idxs := range _reBracketed.FindAllStringSubmatchIndex(template, -1) {
		// Find attribute names to replace and with what handler function to map them to.
		attribute := template[idxs[4]:idxs[5]]
		if fn, ok := custom[attribute]; ok {
			f.handlers = append(f.handlers, fn)
		} else {
			switch attribute {
			case ASC_TIME:
				f.handlers = append(f.handlers, handlerAscTime)
			case FIELDS:
				f.handlers = append(f.handlers, handlerFields)
			case LEVEL_NAME:
				f.handlers = append(f.handlers, handlerLevelName)
			case NAME:
				f.handlers = append(f.handlers, handlerName)
			case MESSAGE:
				f.handlers = append(f.handlers, handlerMessage)
			case PROCESS_PID:
				f.handlers = append(f.handlers, handlerProcessPID)
			case RELATIVE_CREATED:
				f.handlers = append(f.handlers, handlerRelativeCreated)
			case SHORT_LEVEL_NAME:
				f.handlers = append(f.handlers, handlerShortLevelName)
			case HOST:
				f.handlers = append(f.handlers, handlerHost)
			case USER:
				f.handlers = append(f.handlers, handlerUser)
			case AUTH_USER_ID:
				f.handlers = append(f.handlers, handlerAuthUserID)
			case METHOD:
				f.handlers = append(f.handlers, handlerMethod)
			case REQUEST:
				f.handlers = append(f.handlers, handlerRequest)
			case RESPONSE_CODE:
				f.handlers = append(f.handlers, handlerResponseCode)
			case BYTES:
				f.handlers = append(f.handlers, handlerBytes)
			case REF_ERRER:
				f.handlers = append(f.handlers, handlerRefErrer)
			case AGENT:
				f.handlers = append(f.handlers, handlerAgent)
			case USER_IDENTIFIER:
				f.handlers = append(f.handlers, handlerUserIdentifier)
			case DATETIME:
				f.handlers = append(f.handlers, handlerDateTime)
			case TIMESTAMP:
				f.handlers = append(f.handlers, handlerTimeStamp)
			case MODULE:
				f.handlers = append(f.handlers, handlerModule)
			case SEVERITY:
				f.handlers = append(f.handlers, handlerSeverity)
			case THREAD_ID:
				f.handlers = append(f.handlers, handlerThreadID)
			case CLIENT:
				f.handlers = append(f.handlers, handlerClient)
			default:
				continue
			}
		}
		f.attributes[attribute] = true

		// Add segments of the template that do not match regexp (between attributes).
		if segmentsPos < idxs[0] {
			segments = append(segments, template[segmentsPos:idxs[0]])
		}

		// Keep track of padded (y-x > 0) string (== 's') attributes for ANSI color handling.
		if template[idxs[6]:idxs[7]] == "s" && idxs[3]-idxs[2] > 0 {
			start := 0
			for _, s := range segments {
				start += len(s)
			}
			end := start + idxs[3] - idxs[0] + idxs[7] - idxs[6]
			f.handleColors = append(f.handleColors, [...]int{len(f.handlers) - 1, start, end})
		}

		// Update segments.
		segments = append(segments, template[idxs[0]:idxs[3]]+template[idxs[6]:idxs[7]])
		segmentsPos = idxs[1]
	}

	// Add trailing segments of the template that did not match the regexp (newline).
	if segmentsPos < len(template) {
		segments = append(segments, template[segmentsPos:])
	}

	// Join segments.
	f.template = strings.Join(segments, "")
}

func handlerHost(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "Host"), nil
}

func handlerUser(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter,"User"), nil
}

func handlerAuthUserID(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return 12, nil
}

func handlerMethod(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter,"Method"), nil
}

func handlerRequest(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "Request"), nil
}

func handlerResponseCode(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return 200, nil
}

func handlerBytes(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return []byte("Hola mundo"), nil
}

func handlerRefErrer(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "RefErrer"), nil
}

func handlerAgent(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "Agent"), nil
}

func handlerUserIdentifier(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "userIdentifier"), nil
}

func handlerDateTime(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return entry.Time.Format(formatter.timestampFormat), nil
}

func handlerTimeStamp(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return entry.Time.Format(formatter.timestampFormat), nil
}

func handlerModule(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "Module"), nil
}

func handlerSeverity(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "Severity"), nil
}

func handlerThreadID(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return 93, nil
}

func handlerClient(entry *logrus.Entry, formatter *customFormatter) (interface{}, error) {
	return Color(entry, formatter, "Client"), nil
}

