package http

import (
	"time"

	"github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

const (
	DefaultReadTimeout           = 20 * time.Second
	DefaultWriteTimeout          = 20 * time.Second
	DefaultIdleTimeout           = 20 * time.Second
	DefaultReadHeaderTimeout     = 20 * time.Second
	DefaultResponseHeaderTimeout = 20 * time.Second
	DefaultExpectContinueTimeout = 20 * time.Second
	DefaultIdleConnTimeout       = 20 * time.Second
)

type APITimeoutConfig struct {
	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request decisions on each request body's acceptable deadline or upload rate, most users will prefer to use ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout uint64 `mapstructure:"read_timeout,omitempty"`

	// WriteTimeout is the maximum duration before timing out writes of the response. It is reset whenever a new request's header is read. Like ReadTimeout, it does not let Handlers make decisions on a per-request basis.
	WriteTimeout uint64 `mapstructure:"write_timeout,omitempty"`

	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled. If IdleTimeout is zero, the value of ReadTimeout is used. If both are zero, ReadHeaderTimeout is used.
	IdleTimeout uint64 `mapstructure:"idle_timeout,omitempty"`

	// ReadHeaderTimeout is the amount of time allowed to read request headers. The connection's read deadline is reset after reading the headers and the Handler can decide what is considered too slow for the body.
	ReadHeaderTimeout uint64 `mapstructure:"read_header_timeout,omitempty"`

	// IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.
	// Zero means no limit.
	IdleConnTimeout uint64 `mapstructure:"idle_connection_timeout,omitempty"`

	// ResponseHeaderTimeout, if non-zero, specifies the amount of time to wait for a server's response headers after fully writing the request (including its body, if any). This time does not include the time to read the response body.
	ResponseHeaderTimeout uint64 `mapstructure:"response_header_timeout,omitempty"`

	// ExpectContinueTimeout, if non-zero, specifies the amount of time to wait for a server's first response headers after fully writing the request headers if the request has an "Expect: 100-continue" header. Zero means no timeout and causes the body to be sent immediately, without waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout uint64 `mapstructure:"expect_continue_timeout,omitempty"`

	//  timeout time manager
	Dialer *DialerConfig `mapstructure:"dialer,omitempty"`
}

func (a *APITimeoutConfig) Parser() (*http.APITimeout, *errorAVA.Error) {

	var readTimeout time.Duration
	if a.ReadTimeout == 0 {
		readTimeout = DefaultReadTimeout
	} else {
		readTimeout = time.Duration(a.ReadTimeout * uint64(time.Second))
	}

	var writeTimeout time.Duration
	if a.WriteTimeout == 0 {
		writeTimeout = DefaultWriteTimeout
	} else {
		writeTimeout = time.Duration(a.WriteTimeout * uint64(time.Second))
	}

	var idleTimeout time.Duration
	if a.IdleTimeout == 0 {
		idleTimeout = DefaultIdleTimeout
	} else {
		idleTimeout = time.Duration(a.IdleTimeout * uint64(time.Second))
	}

	var readHeaderTimeout time.Duration
	if a.ReadHeaderTimeout == 0 {
		readHeaderTimeout = DefaultReadHeaderTimeout
	} else {
		readHeaderTimeout = time.Duration(a.ReadHeaderTimeout * uint64(time.Second))
	}

	var responseHeaderTimeout time.Duration
	if a.ResponseHeaderTimeout == 0 {
		responseHeaderTimeout = DefaultResponseHeaderTimeout
	} else {
		responseHeaderTimeout = time.Duration(a.ResponseHeaderTimeout * uint64(time.Second))
	}

	var expectContinueTimeout time.Duration
	if a.ExpectContinueTimeout == 0 {
		expectContinueTimeout = DefaultExpectContinueTimeout
	} else {
		expectContinueTimeout = time.Duration(a.ExpectContinueTimeout * uint64(time.Second))
	}

	var idleConnTimeout time.Duration
	if a.IdleConnTimeout == 0 {
		idleConnTimeout = DefaultIdleConnTimeout
	} else {
		idleConnTimeout = time.Duration(a.IdleConnTimeout * uint64(time.Second))
	}

	var dialer *http.Dialer
	var errDialer *errorAVA.Error
	if a.Dialer == nil {
		dialer, errDialer = http.NewDialerDefault()
	} else {
		dialer, errDialer = a.Dialer.Parser()
	}
	if errDialer != nil {
		return nil, errDialer
	}

	return http.NewAPITimeout(
		readTimeout,
		writeTimeout,
		idleTimeout,
		readHeaderTimeout,
		responseHeaderTimeout,
		expectContinueTimeout,
		idleConnTimeout,
		dialer)
}

func (a *APITimeoutConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(a)
}
