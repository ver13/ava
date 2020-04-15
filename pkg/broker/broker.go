package broker

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

var (
	DefaultBroker BrokerI = NewBroker()
)

func Init(opts ...Option) *errorAVA.Error {
	return DefaultBroker.Init(opts...)
}

func Connect() *errorAVA.Error {
	return DefaultBroker.Connect()
}

func Disconnect() *errorAVA.Error {
	return DefaultBroker.Disconnect()
}

func Publish(topic string, msg *Message, opts ...PublishOption) *errorAVA.Error {
	return DefaultBroker.Publish(topic, msg, opts...)
}

func Subscribe(topic string, handler Handler, opts ...SubscribeOption) (SubscriberI, *errorAVA.Error) {
	return DefaultBroker.Subscribe(topic, handler, opts...)
}

func String() string {
	return DefaultBroker.String()
}
