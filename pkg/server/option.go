package server

type Option func(*Options)

// DefaultOptions returns config options for the default service
func DefaultOptions() Options {
	return DefaultServer.Options()
}
