package api

import (
	"github.com/ver13/ava/pkg/registry"
)

// Service represents an API service
type Service struct {
	// Name of service
	Name string
	// The endpoint for this service
	Endpoint *Endpoint
	// Versions of this service
	Services []*registry.Service
}
