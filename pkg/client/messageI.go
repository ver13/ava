package client

// Message is the interface for publishing asynchronously
type MessageI interface {
	Topic() string
	Payload() interface{}
	ContentType() string
}
