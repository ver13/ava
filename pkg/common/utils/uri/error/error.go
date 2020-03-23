package error

const (
	statusInvalidHostCode   = 1
	statusURLParseWrongCode = 2
)

var statusText = map[int]string{
	statusInvalidHostCode:   "Invalid host.",
	statusURLParseWrongCode: "URL parse wrong.",
}

// statusTextFunc returns a text for the General status code. It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
