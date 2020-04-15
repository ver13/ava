package registry

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Watcher is an interface that returns updates
// about services within the registry.
type WatcherI interface {
	// Next is a blocking call
	Next() (*Result, *errorAVA.Error)
	Stop()
}
