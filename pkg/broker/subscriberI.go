package broker

// Subscriber is a convenience return type for the Subscribe method
type SubscriberI interface {
	Options() SubscribeOptions
	Topic() string
	Unsubscribe() error
}
