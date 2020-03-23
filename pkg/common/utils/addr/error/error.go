package error

const (
	StatusFailedGetInterfacesCode  = 1
	StatusIPNotFountCode           = 2
	StatusReadingEtcHostsWrongCode = 3
)

var statusText = map[int]string{
	StatusFailedGetInterfacesCode:  "Failed get interfaces.",
	StatusIPNotFountCode:           "IP not fount.",
	StatusReadingEtcHostsWrongCode: "Reading/ etc/hosts wrong.",
}

// StatusText returns a text for the General status code. It returns the empty string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
