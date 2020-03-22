package error

const (
	statusDeserializer      = 1
	statusSerializer        = 2
	statusNotImplemented    = 3
	statusSerializerUnknown = 4
)

var statusText = map[int]string{
	statusDeserializer:      "Deserializer error.",
	statusSerializer:        "Serializer error.",
	statusNotImplemented:    "Serializer not implemented.",
	statusSerializerUnknown: "Serializer type unknown.",
}

// statusTextFunc returns a text for the encoder status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
