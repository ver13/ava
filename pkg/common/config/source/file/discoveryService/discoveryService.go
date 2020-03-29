package discoveryService

import (
	"context"
	"time"

	httpFileSoutceConfigAVA "github.com/ver13/ava/pkg/common/config/source/file/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type DiscoveryService struct {
	// Set of hosts of the api
	Addrs []string `mapstructure:"addrs"`

	Secure bool `mapstructure:"secure"`

	Timeout time.Duration

	TLS *httpFileSoutceConfigAVA.TLSConfig `mapstructure:"tls,omitempty"`
}

func (b *DiscoveryService) Parser() (*Options, *errorAVA.Error) {
	tls, err := b.TLS.Parser()
	if err != nil {
		return nil, err
	}

	return &Options{
		Addrs:     b.Addrs,
		Secure:    b.Secure,
		Timeout:   b.Timeout,
		TLSConfig: tls,
		Context:   context.Background(),
	}, nil
}

func (b *DiscoveryService) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(b)
}
