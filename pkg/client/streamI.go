package client

import (
	"context"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Stream is the inteface for a bidirectional synchronous stream
type StreamI interface {
	// Context for the stream
	Context() context.Context
	// The request made
	Request() RequestI
	// The response read
	Response() ResponseI
	// Send will encode and send a request
	Send(interface{}) *errorAVA.Error
	// Recv will decode and read a response
	Recv(interface{}) *errorAVA.Error
	// Error returns the stream error
	Error() *errorAVA.Error
	// Close closes the stream
	Close() *errorAVA.Error
}
