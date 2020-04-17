package service

import (
	"github.com/ver13/ava/pkg/client"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	"github.com/ver13/ava/pkg/server"
)

// Service is an interface for a micro service
type ServiceI interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.ClientI
	// Server is for handling requests and events
	Server() server.ServerI
	// Run the service
	Run() *errorAVA.Error
	// The service implementation
	String() string
}
