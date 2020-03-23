package uri

import (
	"sync"
)

var (
	f *uri

	once sync.Once
)

// NewURIParser creates a new URIParser using the package variable RoutingPattern
func init() {
	once.Do(func() {
		f = &uri{RoutingPattern}
	})
}

func GetInstance() *uri {
	return f
}
