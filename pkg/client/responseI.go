package client

import (
	"github.com/ver13/ava/pkg/common/codec"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Response is the response received from a service
type ResponseI interface {
	// Read the response
	Codec() codec.ReaderI
	// read the header
	Header() map[string]string
	// Read the undecoded response
	Read() ([]byte, *errorAVA.Error)
}
