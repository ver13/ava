package transport

import (
	"time"
)

type Message struct {
	Header map[string]string
	Body   []byte
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultTransport TransportI = newHTTPTransport()

	DefaultDialTimeout = time.Second * 5
)

func NewTransport(opts ...Option) TransportI {
	return newHTTPTransport(opts...)
}
