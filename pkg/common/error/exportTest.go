package error

var (
	StatusGroupUnknown    = statusGroupUnknown
	StatusSubgroupUnknown = statusSubgroupUnknown
)

func StatusTextFunc(code int) string {
	return statusTextFunc(code)
}
