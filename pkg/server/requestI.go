package server

import (
	codecAVA "github.com/ver13/ava/pkg/common/codec"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Request is a synchronous request interface
type RequestI interface {
	// Service name requested
	Service() string
	// The action requested
	Method() string
	// Endpoint name requested
	Endpoint() string
	// Content type provided
	ContentType() string
	// Header of the request
	Header() map[string]string
	// Body is the initial decoded value
	Body() interface{}
	// Read the undecoded request body
	Read() ([]byte, *errorAVA.Error)
	// The encoded message stream
	Codec() codecAVA.ReaderI
	// Indicates whether its a stream
	Stream() bool
}
