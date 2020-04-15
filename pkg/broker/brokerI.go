// Package broker is an interface used for asynchronous messaging
package broker

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Broker is an interface used for asynchronous messaging.
type BrokerI interface {
	Init(...Option) *errorAVA.Error
	Options() Options
	Address() string
	Connect() *errorAVA.Error
	Disconnect() *errorAVA.Error
	Publish(topic string, m *Message, opts ...PublishOption) *errorAVA.Error
	Subscribe(topic string, h Handler, opts ...SubscribeOption) (SubscriberI, *errorAVA.Error)
	String() string
}
