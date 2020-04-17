package client

import (
	"context"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	"github.com/ver13/ava/pkg/registry"
)

// CallFunc represents the individual call func
type CallFunc func(ctx context.Context, node *registry.Node, req RequestI, rsp interface{}, opts CallOptions) *errorAVA.Error

// CallWrapper is a low level wrapper for the CallFunc
type CallWrapper func(CallFunc) CallFunc

// Wrapper wraps a client and returns a client
type Wrapper func(ClientI) ClientI

// StreamWrapper wraps a Stream and returns the equivalent
type StreamWrapper func(StreamI) StreamI
