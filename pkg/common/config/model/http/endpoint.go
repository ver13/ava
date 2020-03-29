package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Endpoint struct {
	Method          HTTPVerbType
	CacheTTL        time.Duration
	Timeout         time.Duration
	ConcurrentCalls int
	OutputEncoding  OutputEncodingType
	HeadersToPass   []string
	URL             string
}

func NewEndpointDefault() (*Endpoint, *errorAVA.Error) {
	panic("Not implemented.")
}

func (endpoint *Endpoint) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(endpoint)
}
