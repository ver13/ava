package server

import (
	codecAVA "github.com/ver13/ava/pkg/common/codec"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Response is the response writer for unencoded messages
type ResponseI interface {
	// Encoded writer
	Codec() codecAVA.WriterI
	// Write the header
	WriteHeader(map[string]string)
	// write a response directly to the client
	Write([]byte) *errorAVA.Error
}
