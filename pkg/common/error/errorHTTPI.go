package error

import "net/http"

// AVA HTTP Error Interface
type ErrorHTTPI interface {
	ToJSON() string
	Error() string
	WriteToResponse(w http.ResponseWriter)
}
