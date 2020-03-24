package url

import (
	"sync"
)

var (
	f *URL

	once sync.Once
)

// NewURIParser creates a new URIParser using the package variable RoutingPattern
func init() {
	once.Do(func() {
		f = new(URL)
		f.SetRoutingPattern(RoutingPattern)
	})
}

func GetInstance() *URL {
	return f
}
