package error

const (
	statusLoggerLevelUnknown     = 1
	statusLoggerFormatterUnknown = 2
	statusNotImplemented         = 3
	statusFormatterWrong         = 4
	statusInvalidConfig          = 5
	statusLevelWrong             = 6
	statusNotAvailable           = 7
	statusHandlerError           = 8
)

var statusText = map[int]string{
	statusLoggerLevelUnknown:     "Logger level unknown.",
	statusLoggerFormatterUnknown: "Logger formatter unknown.",
	statusNotImplemented:         "Logger not implemented.",
	statusFormatterWrong:         "Logger formatter wrong.",
	statusInvalidConfig:          "Logger invalid config.",
	statusLevelWrong:             "Logger level wrong.",
	statusNotAvailable:           "Logger not available.",
	statusHandlerError:           "Logger handler error.",
}

// StatusText returns a text for the logger status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
