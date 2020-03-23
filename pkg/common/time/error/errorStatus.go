package error

const (
	statusTimeout         = 1
	statusLoadLocation    = 2
	statusTimeFormat      = 3
	statusParseInLocation = 4
)

var statusText = map[int]string{
	statusTimeout:         "Timeout.",
	statusLoadLocation:    "Load location time error.",
	statusTimeFormat:      "Time format error.",
	statusParseInLocation: "Parser in location error.",
}

// statusTextFunc returns a text for the General status code. It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
