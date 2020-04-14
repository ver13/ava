package registry

import (
	"time"
)

// Event is registry event
type Event struct {
	// Id is registry id
	Id string
	// Type defines type of event
	Type EventType
	// Timestamp is event timestamp
	Timestamp time.Time
	// Service is registry service
	Service *Service
}
