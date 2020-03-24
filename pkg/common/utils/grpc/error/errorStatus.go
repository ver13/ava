package error

const (
	statusMethodNameWrong  = 1
)

var statusText = map[int]string{
	statusMethodNameWrong:  "Method name wrong.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
