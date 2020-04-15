package server

import codecAVA "github.com/ver13/ava/pkg/common/codec"

// Message is an async message interface
type MessageI interface {
	// Topic of the message
	Topic() string
	// The decoded payload value
	Payload() interface{}
	// The content type of the payload
	ContentType() string
	// The raw headers of the message
	Header() map[string]string
	// The raw body of the message
	Body() []byte
	// Codec used to decode the message
	Codec() codecAVA.ReaderI
}
