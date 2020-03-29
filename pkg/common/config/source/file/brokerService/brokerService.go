package brokerService

import (
	"context"

	httpFileSoutceConfigAVA "github.com/ver13/ava/pkg/common/config/source/file/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type BrokerService struct {
	// Set of hosts of the api
	Addrs []string `mapstructure:"addrs"`

	Secure bool `mapstructure:"secure"`

	Codec string `mapstructure:"messageCoder"`

	TLS *httpFileSoutceConfigAVA.TLSConfig `mapstructure:"tls,omitempty"`
}

func (b *BrokerService) Parser() (*Options, *errorAVA.Error) {
	tls, err := b.TLS.Parser()
	if err != nil {
		return nil, err
	}

	return &Options{
		Addrs:        b.Addrs,
		Secure:       b.Secure,
		MessageCoder: nil,
		TLSConfig:    tls,
		Context:      context.Background(),
	}, nil
}

func (b *BrokerService) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(b)
}
