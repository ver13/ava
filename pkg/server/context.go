package server

import (
	"context"
	"sync"
)

type serverKey struct{}

func wait(ctx context.Context) *sync.WaitGroup {
	if ctx == nil {
		return nil
	}
	wg, ok := ctx.Value("wait").(*sync.WaitGroup)
	if !ok {
		return nil
	}
	return wg
}

func FromContext(ctx context.Context) (ServerI, bool) {
	c, ok := ctx.Value(serverKey{}).(ServerI)
	return c, ok
}

func NewContext(ctx context.Context, s ServerI) context.Context {
	return context.WithValue(ctx, serverKey{}, s)
}
