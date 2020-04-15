package transport

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type ListenerI interface {
	Addr() string
	Close() *errorAVA.Error
	Accept(func(SocketI)) *errorAVA.Error
}
