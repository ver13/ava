package error

const (
	statusInvalidHost   = 1
	statusURLParseWrong = 2
)

var statusText = map[int]string{
	statusInvalidHost:   "Invalid host.",
	statusURLParseWrong: "URL parse wrong.",
}

// statusTextFunc returns a text for the General status code. It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
