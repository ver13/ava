package pool

import (
	"time"

	transportGmf "github.com/ValentinEncinasRojas/ava/pkg/transport"
)

type Options struct {
	Transport transportGmf.TransportI
	TTL       time.Duration
	Size      int
}

type OptionFunc func(*Options)

func Size(i int) OptionFunc {
	return func(o *Options) {
		o.Size = i
	}
}

func Transport(t transportGmf.TransportI) OptionFunc {
	return func(o *Options) {
		o.Transport = t
	}
}

func TTL(t time.Duration) OptionFunc {
	return func(o *Options) {
		o.TTL = t
	}
}
