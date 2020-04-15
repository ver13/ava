package server

import (
	"context"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Stream represents a stream established with a client.
// A stream can be bidirectional which is indicated by the request.
// The last error will be left in Error().
// EOF indicates end of the stream.
type StreamI interface {
	Context() context.Context
	Request() RequestI
	Send(interface{}) *errorAVA.Error
	Recv(interface{}) *errorAVA.Error
	Error() *errorAVA.Error
	Close() *errorAVA.Error
}
