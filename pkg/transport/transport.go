package transport

import (
	"time"

	httpTransportAVA "github.com/ver13/ava/pkg/transport/http"
)

type Message struct {
	Header map[string]string
	Body   []byte
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultTransport TransportI = httpTransportAVA.NewTransport()

	DefaultDialTimeout = time.Second * 5
)

func NewTransport(opts ...Option) TransportI {
	return httpTransportAVA.NewTransport(opts...)
}
