package error

// AVA errors status codes
const (
	statusGroupUnknown    = 1
	statusSubgroupUnknown = 2
	serializerJSONCode    = 3
)

var statusText = map[int]string{
	statusGroupUnknown:    "Group type unknown.",
	statusSubgroupUnknown: "Subgroup type unknown.",
	serializerJSONCode:    "Serializer JSON errors.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
