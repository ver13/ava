package client

import (
	"context"
)

// Router manages request routing
type Router interface {
	SendRequest(context.Context, RequestI) (ResponseI, error)
}
