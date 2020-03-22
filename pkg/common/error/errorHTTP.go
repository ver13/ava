package error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorHTTP contains custom code, errors message, and HTTP status code.
type ErrorHTTP struct {
	HTTPStatus int    `json:"status"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

// Error return response for the Error HTTP.
func (e *ErrorHTTP) Error() string {
	return fmt.Sprintf("HTTPStatus: %v, Code: %v, Message: %q",
		e.HTTPStatus, e.Code, e.Message)
}

// ToJSON returns JSON string for a ErrorHTTP.
func (e *ErrorHTTP) ToJSON() string {
	j, err := json.Marshal(e)
	if err != nil {
		return `{"code":50099,"message":"ScrapError.JSONStr: json.Marshal() failed"}`
	}
	return string(j)
}

// WriteToResponse writes response for the Error HTTP.
func (e *ErrorHTTP) WriteToResponse(w http.ResponseWriter) {
	w.WriteHeader(e.HTTPStatus)
	fmt.Fprintf(w, "%s", e.ToJSON())
	// TODO: store e.ToJSON() to ElasticSearch for future analysis
}
