package http

import (
	"net/url"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Backend struct {
	// the name of the group the response should be moved to. If empty, the response is not changed
	Group string
	// HTTP Method of the request to send to the backend
	Method HTTPVerbType
	// Set of hosts of the api
	Host []string
	// False if the hostname should be sanitized
	HostSanitizationDisabled bool
	// URL pattern to use to locate the resource to be consumed
	URL *url.URL
	// set of response fields to remove. If empty, the filter id not used
	Blacklist []string
	// set of response fields to allow. If empty, the filter id not used
	Whitelist []string
	// map of response fields to be renamed and their new names
	Mapping map[string]string
	// the error format
	Encoding OutputEncodingType
	// the response to process is a collection
	IsCollection bool
	// name of the field to extract to the root. If empty, the formater will do nothing
	Target string
	// name of the Discovery Service driver to use
	DiscoveryService string

	// list of keys to be replaced in the url
	URLKeys []string
	// number of concurrent calls this endpoint must send to the api
	ConcurrentCalls int
	// timeout of this backend
	Timeout time.Duration
}

func NewBackendDefault() (*Backend, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewBackend(group string, method HTTPVerbType, host []string, disabled bool, url string, blacklist []string, whitelist []string, mapping map[string]string, encoding OutputEncodingType, collection bool, target string, service string, keys []string, calls int, timeout time.Duration) (*Backend, *errorAVA.Error) {
	panic("Not implemented.")
}

func (backend *Backend) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(backend)
}
