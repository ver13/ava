package error

const (
	statusWriteBinary = 1
	statusTypeUnknown = 2
)

var statusText = map[int]string{
	statusWriteBinary: "Write binary error.",
	statusTypeUnknown: "Type unknown.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
