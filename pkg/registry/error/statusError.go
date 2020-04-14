package error

const (
	statusDeserializer     = 1
	statusMessage          = 2
	statusNotImplemented   = 3
	statusEventTypeUnknown = 4
)

var statusText = map[int]string{
	statusDeserializer:     "Deserializer error.",
	statusMessage:          "Message error.",
	statusNotImplemented:   "Message not implemented.",
	statusEventTypeUnknown: "Event type unknown.",
}

// statusTextFunc returns a text for the encoder status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
