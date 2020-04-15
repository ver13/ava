package server

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Server is a simple server abstraction
type ServerI interface {
	Options() Options
	Init(...Option) *errorAVA.Error
	Handle(HandlerI) *errorAVA.Error
	NewHandler(interface{}, ...HandlerOption) HandlerI
	NewSubscriber(string, interface{}, ...SubscriberOption) SubscriberI
	Subscribe(SubscriberI) *errorAVA.Error
	Start() *errorAVA.Error
	Stop() *errorAVA.Error
	String() string
}
