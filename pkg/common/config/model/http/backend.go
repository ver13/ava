package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Backend struct {
}

func NewBackend(group string, method HTTPVerbType, host []string, disabled bool, url string, blacklist []string, whitelist []string, mapping map[string]string, encoding OutputEncodingType, collection bool, target string, service string, keys []string, calls int, timeout time.Duration) (*Backend, *errorAVA.Error) {
	panic("Not implemented.")
}

func (backend *Backend) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(backend)
}
