package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Dialer struct {
	// DialerTimeout is the maximum amount of time a dial will wait for a connect to complete. If Deadline is also set, it may fail earlier.
	// The default is no timeout.
	// When using TCP and dialing a host name with multiple IP addresses, the timeout may be divided between them.
	// With or without a timeout, the operating system may impose its own earlier timeout. For instance, TCP timeouts are often around 3 minutes.
	DialerTimeout uint64
	// FallbackDelay specifies the length of time to wait before spawning a fallback connection, when DualStack is enabled.
	// If zero, a default delay of 300ms is used.
	DialerFallbackDelay uint64
	// KeepAlive specifies the keep-alive period for an active network connection.
	// If zero, keep-alives are not enabled. Network protocols that do not support keep-alives ignore this field.
	DialerKeepAlive uint64
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
