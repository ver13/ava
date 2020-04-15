package transport

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Transport is an interface which is used for communication between
// services. It uses connection based socket send/recv semantics and
// has various implementations; http, grpc, quic.
type TransportI interface {
	Init(...Option) *errorAVA.Error
	Options() Options
	Dial(addr string, opts ...DialOption) (ClientI, *errorAVA.Error)
	Listen(addr string, opts ...ListenOption) (ListenerI, *errorAVA.Error)
	String() string
}
