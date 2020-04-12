package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type APITimeout struct {
	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	// Because ReadTimeout does not let Handlers make per-request decisions on each request body's acceptable deadline or upload rate, most users will prefer to use ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout uint64

	// WriteTimeout is the maximum duration before timing out writes of the response. It is reset whenever a new request's header is read. Like ReadTimeout, it does not let Handlers make decisions on a per-request basis.
	WriteTimeout uint64

	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled. If IdleTimeout is zero, the value of ReadTimeout is used. If both are zero, ReadHeaderTimeout is used.
	IdleTimeout uint64

	// ReadHeaderTimeout is the amount of time allowed to read request headers. The connection's read deadline is reset after reading the headers and the Handler can decide what is considered too slow for the body.
	ReadHeaderTimeout uint64

	// IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.
	// Zero means no limit.
	IdleConnTimeout uint64

	// ResponseHeaderTimeout, if non-zero, specifies the amount of time to wait for a server's response headers after fully writing the request (including its body, if any). This time does not include the time to read the response body.
	ResponseHeaderTimeout uint64

	// ExpectContinueTimeout, if non-zero, specifies the amount of time to wait for a server's first response headers after fully writing the request headers if the request has an "Expect: 100-continue" header. Zero means no timeout and causes the body to be sent immediately, without waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout uint64

	//  timeout time manager
	Dialer *Dialer
}

func NewAPITimeout(timeout time.Duration, timeout2 time.Duration, timeout3 time.Duration, timeout4 time.Duration, timeout5 time.Duration, timeout6 time.Duration, timeout7 time.Duration, dialer *Dialer) (*APITimeout, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewAPITimeoutDefault() (*APITimeout, *errorAVA.Error) {
	panic("Not implemented.")
}

func (apiTimeout *APITimeout) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(apiTimeout)
}
