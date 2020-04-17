package api

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type APII interface {
	// Register a http handler
	Register(*Endpoint) *errorAVA.Error
	// Register a route
	Deregister(*Endpoint) *errorAVA.Error
	// Init initialises the command line.
	// It also parses further options.
	//Init(...Option) error
	// Options
	//Options() Options
	// String
	String() string
}
