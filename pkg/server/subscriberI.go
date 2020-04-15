package server

import (
	"github.com/ver13/ava/pkg/registry"
)

// Subscriber interface represents a subscription to a given topic using
// a specific subscriber function or object with endpoints.
type SubscriberI interface {
	Topic() string
	Subscriber() interface{}
	Endpoints() []*registry.Endpoint
	Options() SubscriberOptions
}
