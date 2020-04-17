package error

const (
	statusParseFileWrong      = 1
	statusNoDocOnEnum         = 2
	statusWritingHeaderFailed = 3
	statusWritingDataFailed   = 4
	statusFormattingError     = 5
)

var statusText = map[int]string{
	statusParseFileWrong:      "Parse file wrong.",
	statusNoDocOnEnum:         "No doc on enum.",
	statusWritingHeaderFailed: "Writing header failed.",
	statusWritingDataFailed:   "Writing data failed.",
	statusFormattingError:     "Formatting error.",
}

// statusTextFunc returns a text for the encoder status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
