package registry

// Watcher is an interface that returns updates
// about services within the registry.
type WatcherI interface {
	// Next is a blocking call
	Next() (*Result, error)
	Stop()
}
