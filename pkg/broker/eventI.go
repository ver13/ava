package broker

// Event is given to a subscription handler for processing
type EventI interface {
	Topic() string
	Message() *Message
	Ack() error
	Error() error
}
