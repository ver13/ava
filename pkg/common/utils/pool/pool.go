package pool

import (
	"time"

	errorGmf "github.com/ValentinEncinasRojas/ava/errors"

	transportGmf "github.com/ValentinEncinasRojas/ava/pkg/transport"
)

// PoolI is an interface for connection pooling
type PoolI interface {
	// Close the pool
	Close() errorGmf.ErrorGmfI
	// Get a connection
	Get(addr string, opts ...transportGmf.DialOptionFunc) (ConnI, errorGmf.ErrorGmfI)
	// Releaes the connection
	Release(c ConnI, status errorGmf.ErrorGmfI) errorGmf.ErrorGmfI
}

type ConnI interface {
	// unique id of connection
	Id() string
	// time it was created
	Created() time.Time
	// embedded connection
	transportGmf.ClientI
}

func NewPool(opts ...Option) PoolI {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return newPool(options)
}
