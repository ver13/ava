package http

import (
	"fmt"
	"time"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	"github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

const (
	ConcurrentCallsDefault = 1
	DefaultCacheTTL        = 6 * time.Hour
)

type EndpointConfig struct {
	// url pattern to be registered and exposed to the world
	URL string `mapstructure:"url"`
	// HTTP Method of the endpoint (GET, POST, PUT, etc)
	Method string `mapstructure:"Method"`
	// set of definitions of the backends to be linked to this endpoint
	Backend []*BackendConfig `mapstructure:"backend"`
	//
	DiscoveryService string `mapstructure:"discovery_service_name"`
	// number of concurrent calls this endpoint must send to the backends
	ConcurrentCalls int `mapstructure:"concurrent_calls"`
	// timeout of this endpoint
	Timeout time.Duration `mapstructure:"timeout"`
	// duration of the cache header
	CacheTTL time.Duration `mapstructure:"cache_ttl"`
	// list of query string params to be extracted from the URI
	QueryString []string `mapstructure:"querystring_params"`
	// headersToPass defines the list of headers to pass to the backends
	HeadersToPass []string `mapstructure:"headers_to_pass"`
	// OutputEncodingType defines the error strategy to use for the endpoint responses
	OutputEncoding string `mapstructure:"output_encoding"`
}

func NewEndpointConfig(URL string, method string, backend []*BackendConfig, discoveryService string, concurrentCalls int, timeout time.Duration, cacheTTL time.Duration, queryString []string, headersToPass []string, outputEncoding string) (*http.Endpoint, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewEndpointConfigDefault() (*http.Endpoint, *errorAVA.Error) {
	panic("Not implemented.")
}

func (e *EndpointConfig) ReadLocal(fileName string) (*http.Endpoint, *errorAVA.Error) {
	panic("Not implemented.")
}

func (e *EndpointConfig) Parser(api *APIConfig) (*http.Endpoint, *errorAVA.Error) {
	if api == nil {
		return nil, errorConfigAVA.APIIsEmpty(nil, fmt.Sprintf("%p", e))
	}

	endpoint, err := e.initEndpointDefaults(api)
	if err != nil {
		return nil, err
	}

	if e.URL == "" {
		return nil, errorConfigAVA.URLIsEmpty(nil, fmt.Sprintf("%p", e))
	}
	endpoint.URL = e.URL

	endpoint.HeadersToPass = e.HeadersToPass

	concurrentCalls := e.ConcurrentCalls
	if concurrentCalls == 0 {
		concurrentCalls = ConcurrentCallsDefault
	}
	endpoint.ConcurrentCalls = concurrentCalls

	endpoint.Timeout = e.Timeout
	if e.Timeout == 0 {
		endpoint.Timeout = DefaultTimeout
	}

	cacheTTL := e.CacheTTL
	if cacheTTL == 0 {
		cacheTTL = DefaultCacheTTL
	}
	endpoint.CacheTTL = cacheTTL

	var outputEncoding http.OutputEncodingType
	if e.OutputEncoding == "" {
		outputEncoding = http.OutputEncodingTypeJSON
	} else {
		var err error
		outputEncoding, err = http.ParseOutputEncodingType(e.OutputEncoding)
		if err != nil {
			return nil, errorConfigAVA.OutputEncodingWrong(err, fmt.Sprintf("output error: %s", e.OutputEncoding))
		}
	}
	endpoint.OutputEncoding = outputEncoding

	return endpoint, nil
}

func (e *EndpointConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(e)
}

func (e *EndpointConfig) initEndpointDefaults(api *APIConfig) (*http.Endpoint, *errorAVA.Error) {
	endpoint, _ := http.NewEndpointDefault()
	if e.Method == "" {
		endpoint.Method = http.HTTPVerbTypeGET
	} else {
		var err error
		endpoint.Method, err = http.ParseHTTPVerbType(e.Method)
		if err != nil {
			return nil, errorConfigAVA.AllowMethodWrong(err, fmt.Sprintf("Method: %s", e.Method))
		}
	}

	if api.CacheTTL != 0 && e.CacheTTL == 0 {
		endpoint.CacheTTL = api.CacheTTL
	}

	if api.Timeout != 0 && e.Timeout == 0 {
		endpoint.Timeout = api.Timeout
	}

	if e.ConcurrentCalls == 0 {
		endpoint.ConcurrentCalls = 1
	}

	if e.OutputEncoding == "" {
		if api.OutputEncoding != "" {
			var err error
			endpoint.OutputEncoding, err = http.ParseOutputEncodingType(api.OutputEncoding)
			if err != nil {
				return nil, errorConfigAVA.OutputEncodingWrong(err, fmt.Sprintf("encoding: %s", api.OutputEncoding))
			}
		} else {
			endpoint.OutputEncoding = http.OutputEncodingTypeJSON
		}
	}

	return endpoint, nil
}
