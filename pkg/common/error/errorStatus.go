package error

// AVA errors status codes
const (
	GroupUnknownCode    = 1 //
	SubgroupUnknownCode = 2 //
	SerializerJSONCode  = 3 //
)

var statusText = map[int]string{
	GroupUnknownCode:    "Group type unknown.",
	SubgroupUnknownCode: "Subgroup type unknown.",
	SerializerJSONCode:  "Serializer JSON errors.",
}

// StatusText returns a text for the General status code.
// It returns the empty string if the code is unknown.
func StatusTextFunc(code int) string {
	return statusText[code]
}
