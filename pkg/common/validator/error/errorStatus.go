package error

const (
	statusInvalidData  = 1
	statusNotFoundData = 2
)

var statusText = map[int]string{
	statusInvalidData:  "Invalid data.",
	statusNotFoundData: "Not found data.",
}

// statusTextFunc returns a text for the General status code. It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
