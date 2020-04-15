package server

import (
	"github.com/ver13/ava/pkg/registry"
)

// Handler interface represents a request handler. It's generated
// by passing any type of public concrete object with endpoints into server.NewHandler.
type HandlerI interface {
	Name() string
	Handler() interface{}
	Endpoints() []*registry.Endpoint
	Options() HandlerOptions
}
