package broker

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(EventI) *errorAVA.Error
