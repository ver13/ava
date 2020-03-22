package general

// AVA errors status codes
const (
	NotEqualCode       = 1
	IsNilCode          = 2
	DeepCopyWrongCode  = 3
	UnknownErrorCode   = 4
	SerializerJSONCode = 7
)

var statusText = map[int]string{
	NotEqualCode:      "Error not equal.",
	IsNilCode:         "ErrorHTTP is NIL.",
	DeepCopyWrongCode: "Deep copy wrong.",
	UnknownErrorCode:  "Unknown error.",
}

// StatusText returns a text for the General status code.
// It returns the empty string if the code is unknown.
func StatusTextFunc(code int) string {
	return statusText[code]
}
