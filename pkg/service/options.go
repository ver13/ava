package service

import (
	"context"
	"time"

	"github.com/ver13/ava/pkg/broker"
	"github.com/ver13/ava/pkg/client"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	"github.com/ver13/ava/pkg/registry"
	"github.com/ver13/ava/pkg/server"
	"github.com/ver13/ava/pkg/transport"
)

type Options struct {
	Broker    broker.BrokerI
	Client    client.ClientI
	Server    server.ServerI
	Registry  registry.RegistryI
	Transport transport.TransportI

	// Before and After funcs
	BeforeStart []func() *errorAVA.Error
	BeforeStop  []func() *errorAVA.Error
	AfterStart  []func() *errorAVA.Error
	AfterStop   []func() *errorAVA.Error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	opt := Options{
		Broker:    broker.DefaultBroker,
		Client:    client.DefaultClient,
		Server:    server.DefaultServer,
		Registry:  registry.DefaultRegistry,
		Transport: transport.DefaultTransport,
		Context:   context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Broker(b broker.BrokerI) Option {
	return func(o *Options) {
		o.Broker = b
		// Update Client and Server
		o.Client.Init(client.Broker(b))
		o.Server.Init(server.Broker(b))
	}
}

func Client(c client.ClientI) Option {
	return func(o *Options) {
		o.Client = c
	}
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service.
// Can be used for extra option values.
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func Server(s server.ServerI) Option {
	return func(o *Options) {
		o.Server = s
	}
}

// Registry sets the registry for the service
// and the underlying components
func Registry(r registry.RegistryI) Option {
	return func(o *Options) {
		o.Registry = r
		// Update Client and Server
		o.Client.Init(client.Registry(r))
		o.Server.Init(server.Registry(r))
		// Update Broker
		o.Broker.Init(broker.Registry(r))
	}
}

// Transport sets the transport for the service
// and the underlying components
func Transport(t transport.TransportI) Option {
	return func(o *Options) {
		o.Transport = t
		// Update Client and Server
		o.Client.Init(client.Transport(t))
		o.Server.Init(server.Transport(t))
	}
}

// Convenience options

// Address sets the address of the server
func Address(addr string) Option {
	return func(o *Options) {
		o.Server.Init(server.Address(addr))
	}
}

// Name of the service
func Name(n string) Option {
	return func(o *Options) {
		o.Server.Init(server.Name(n))
	}
}

// Version of the service
func Version(v string) Option {
	return func(o *Options) {
		o.Server.Init(server.Version(v))
	}
}

// Metadata associated with the service
func Metadata(md map[string]string) Option {
	return func(o *Options) {
		o.Server.Init(server.Metadata(md))
	}
}

// RegisterTTL specifies the TTL to use when registering the service
func RegisterTTL(t time.Duration) Option {
	return func(o *Options) {
		o.Server.Init(server.RegisterTTL(t))
	}
}

// RegisterInterval specifies the interval on which to re-register
func RegisterInterval(t time.Duration) Option {
	return func(o *Options) {
		o.Server.Init(server.RegisterInterval(t))
	}
}

// Before and Afters

func BeforeStart(fn func() *errorAVA.Error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func BeforeStop(fn func() *errorAVA.Error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

func AfterStart(fn func() *errorAVA.Error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func AfterStop(fn func() *errorAVA.Error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
