package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Dialer struct {
}

func NewDialer(timeout time.Duration, delay time.Duration, alive time.Duration) (*Dialer, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewDialerDefault() (*Dialer, *errorAVA.Error) {
	panic("Not implemented.")
}

func (dialer *Dialer) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(dialer)
}
