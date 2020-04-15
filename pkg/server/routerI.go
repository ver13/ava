package server

import (
	"context"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Router handle serving messages
type RouterI interface {
	// ProcessMessage processes a message
	ProcessMessage(context.Context, MessageI) *errorAVA.Error
	// ServeRequest processes a request to completion
	ServeRequest(context.Context, RequestI, ResponseI)
}
