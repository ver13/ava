package error

const (
	statusFailedGetInterfaces  = 1
	statusIPNotFount           = 2
	statusReadingEtcHostsWrong = 3
)

var statusText = map[int]string{
	statusFailedGetInterfaces:  "Failed get interfaces.",
	statusIPNotFount:           "IP not fount.",
	statusReadingEtcHostsWrong: "Reading/ etc/hosts wrong.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
