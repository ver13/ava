package error

const (
	statusHostPortWrong = 1
	statusPortWrong     = 2
	statusAddrWrong     = 3
)

var statusText = map[int]string{
	statusHostPortWrong: "Host port wrong.",
	statusPortWrong:     "Port wrong.",
	statusAddrWrong:     "Addr wrong.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
