package http

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type API struct {
	// set of endpoint definitions
	Endpoints []*Endpoint
	// default TTL for GET
	CacheTTL time.Duration
	// default set of hosts
	Host []string
	//
	Port uint64
	// version code of the configurationServiceI
	Version int
	// OutputEncodingType defines the default error strategy to use for the endpoint responses
	OutputEncoding string
	// defafult timeout
	Timeout time.Duration

	// api time manager
	APITime *APITimeout

	CORS *CORS

	// disableKeepAlives, if true, prevents re-use of TCP connections between different HTTP requests.
	DisableKeepAlives bool
	// disableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-encoding: gzip"
	// request header when the Request contains no existing
	// Accept-encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression bool
	// maxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int
	// maxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int

	// disableStrictREST flags if the REST enforcement is disabled
	DisableStrictREST bool

	// run AVA in debug mode
	Debug bool
}

func NewAPIDefault() (*API, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewAPI(endpoints []*Endpoint, ttl time.Duration, host []string, port uint64, version int, encoding OutputEncodingType, timeout time.Duration, config *APITimeout, cors *CORS, alives bool, compression bool, conns int, host2 int, rest bool, debug bool) (*API, *errorAVA.Error) {
	panic("Not implemented.")
}

func (api *API) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(api)
}
