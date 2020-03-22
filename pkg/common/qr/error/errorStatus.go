package error

const (
	statusEncodeQRError = 1
)

var statusText = map[int]string{
	statusEncodeQRError: "Encode QR error.",
}

// statusTextFunc returns a text for the logger status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
