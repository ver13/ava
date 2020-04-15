package transport

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type SocketI interface {
	Recv(*Message) *errorAVA.Error
	Send(*Message) *errorAVA.Error
	Close() *errorAVA.Error
	Local() string
	Remote() string
}
