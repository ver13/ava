package http

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	stringAVA "github.com/ver13/ava/pkg/common/string"
	urlUtilsAVA "github.com/ver13/ava/pkg/common/utils/url"
)

const (
	// BracketsRouterPatternBuilder uses brackets as route params delimiter
	BracketsRouterPatternBuilder = iota
	// ColonRouterPatternBuilder use a colon as route param delimiter
	ColonRouterPatternBuilder
	// DefaultMaxIdleConnsPerHost is the default value for the maxIdleConnsPerHost param
	DefaultMaxIdleConnsPerHost = 250
	// DefaultTimeout is the default value to use for the ServiceConfig.timeout param
	DefaultTimeout = 2 * time.Second

	// ConfigVersion is the current version of the error struct
	ConfigVersion = 1
)

var (
	simpleURLKeysPattern   = regexp.MustCompile(`{([a-zA-Z\-_0-9]+)}`)
	debugPattern           = "^[^/]|/__debug(/.*)?$"
	errInvalidHost         = errors.New("invalid host")
	errInvalidNoOpEncoding = errors.New("can not use NoOp error with more than one backends connected to the same endpoint")
	defaultPort            = 8080
)

// APIConfig defines the AVA stackMicroservices
type APIConfig struct {
	// set of endpoint definitions
	Endpoints []*EndpointConfig `mapstructure:"endpoints"`
	// default TTL for GET
	CacheTTL time.Duration `mapstructure:"cache_ttl"`
	// default set of hosts
	Host []string `mapstructure:"host"`
	//
	Port uint64 `mapstructure:"Port"`
	// version code of the configurationServiceI
	Version int `mapstructure:"version"`
	// OutputEncodingType defines the default error strategy to use for the endpoint responses
	OutputEncoding string `mapstructure:"output_encoding"`
	// defafult timeout
	Timeout time.Duration `mapstructure:"timeout"`

	// api time manager
	APITime *APITimeoutConfig `mapstructure:"time"`

	CORS *CORSConfig `mapstructure:"cors"`

	// disableKeepAlives, if true, prevents re-use of TCP connections between different HTTP requests.
	DisableKeepAlives bool `mapstructure:"disable_keep_alives"`
	// disableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-encoding: gzip"
	// request header when the Request contains no existing
	// Accept-encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression bool `mapstructure:"disable_compression"`
	// maxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int `mapstructure:"max_idle_connections"`
	// maxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int `mapstructure:"max_idle_connections_per_host"`

	// disableStrictREST flags if the REST enforcement is disabled
	DisableStrictREST bool `mapstructure:"disable_rest"`

	// run AVA in debug mode
	Debug bool `mapstructure:"debug"`
}

func NewAPIConfig(endpoints []*EndpointConfig, cacheTTL time.Duration, host []string, port uint64, version int, outputEncoding string, timeout time.Duration, APITime *APITimeoutConfig, CORS *CORSConfig, disableKeepAlives bool, disableCompression bool, maxIdleConns int, maxIdleConnsPerHost int,
	disableStrictREST bool, debug bool) (*httpModelConfigAVA.API, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewAPIConfigDefault() (*httpModelConfigAVA.API, *errorAVA.Error) {
	api := &httpModelConfigAVA.API{
		Endpoints:           nil,
		CacheTTL:            0,
		Host:                nil,
		Port:                0,
		Version:             0,
		OutputEncoding:      "",
		Timeout:             0,
		APITime:             nil,
		CORS:                nil,
		DisableKeepAlives:   false,
		DisableCompression:  false,
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 0,
		DisableStrictREST:   false,
		Debug:               false,
	}
	return api, nil
}

func (a *APIConfig) ReadLocal(fileName string) (*httpModelConfigAVA.API, *errorAVA.Error) {
	panic("Not implemented.")
}

func (a *APIConfig) Parser() (*httpModelConfigAVA.API, *errorAVA.Error) {
	if len(a.Host) == 0 {
		a.Host = []string{"localhost"}
	}
	a.Host = urlUtilsAVA.CleanHosts(a.Host)

	if a.Version != ConfigVersion {
		return nil, errorConfigAVA.ConfigVersionWrong(nil, fmt.Sprintf("Unsupported version: %d (want: %d)", a.Version, ConfigVersion))
	}

	var endpoints []*httpModelConfigAVA.Endpoint
	for _, e := range a.Endpoints {
		endpoint, err := e.Parser(a)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	outputEncoding, err := httpModelConfigAVA.ParseOutputEncodingType(stringAVA.StringToUpper(a.OutputEncoding))
	if err != nil {
		return nil, errorConfigAVA.OutputEncodingWrong(err, fmt.Sprintf("output error: %s", a.OutputEncoding))
	}

	var time *httpModelConfigAVA.APITimeout
	var errTime *errorAVA.Error
	if a.APITime == nil {
		time, errTime = httpModelConfigAVA.NewAPITimeoutDefault()
	} else {
		time, errTime = a.APITime.Parser()
	}
	if errTime != nil {
		time, _ = httpModelConfigAVA.NewAPITimeoutDefault()
	}

	var cors *httpModelConfigAVA.CORS
	var errCors *errorAVA.Error
	if a.CORS == nil {
		cors, errCors = httpModelConfigAVA.NewCORSDefault()
	} else {
		cors, errCors = a.CORS.Parser()
	}
	if errCors != nil {
		return nil, errCors
	}

	return httpModelConfigAVA.NewAPI(endpoints,
		a.CacheTTL,
		a.Host,
		a.Port,
		a.Version,
		outputEncoding,
		a.Timeout,
		time,
		cors,
		a.DisableKeepAlives,
		a.DisableCompression,
		a.MaxIdleConns,
		a.MaxIdleConnsPerHost,
		a.DisableStrictREST,
		a.Debug)
}

func (a *APIConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(a)
}
