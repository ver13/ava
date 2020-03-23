package error

const (
	StatusHostPortWrongCode = 1
	StatusPortWrongCode     = 2
	StatusAddrWrongCode     = 3
)

var statusText = map[int]string{
	StatusHostPortWrongCode: "Host port wrong.",
	StatusPortWrongCode:     "Port wrong.",
	StatusAddrWrongCode:     "Addr wrong.",
}

// StatusText returns a text for the General status code. It returns the empty string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
