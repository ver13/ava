package error

const (
	statusFlushTokenXml  = 1
	statusEncodeTokenXml = 2
	statusTimeParser     = 3
)

var statusText = map[int]string{
	statusFlushTokenXml:  "Flush token xml error.",
	statusEncodeTokenXml: "Encode token xml error.",
	statusTimeParser:     "apiTime parser error.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
