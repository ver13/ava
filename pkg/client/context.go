package client

import (
	"context"
)

type clientKey struct{}

func FromContext(ctx context.Context) (ClientI, bool) {
	c, ok := ctx.Value(clientKey{}).(ClientI)
	return c, ok
}

func NewContext(ctx context.Context, c ClientI) context.Context {
	return context.WithValue(ctx, clientKey{}, c)
}
