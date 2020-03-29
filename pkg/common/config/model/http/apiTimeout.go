package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type APITimeout struct {
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
