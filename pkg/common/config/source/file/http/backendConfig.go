package http

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorGeneralAVA "github.com/ver13/ava/pkg/common/error/general"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	uriUtilsAVA "github.com/ver13/ava/pkg/common/utils/url"
)

// BackendConfig defines how AVA should connect to the BackendConfig stackMicroservices (the api resource to consume) and how it should process the received response
type BackendConfig struct {
	// the name of the group the response should be moved to. If empty, the response is not changed
	Group string `mapstructure:"group"`
	// HTTP Method of the request to send to the backend
	Method string `mapstructure:"Method"`
	// Set of hosts of the api
	Host []string `mapstructure:"host"`
	// False if the hostname should be sanitized
	HostSanitizationDisabled bool `mapstructure:"disable_host_sanitize"`
	// url pattern to use to locate the resource to be consumed
	Url string `mapstructure:"url_pattern"`
	// set of response fields to remove. If empty, the filter id not used
	Blacklist []string `mapstructure:"blacklist"`
	// set of response fields to allow. If empty, the filter id not used
	Whitelist []string `mapstructure:"whitelist"`
	// map of response fields to be renamed and their new names
	Mapping map[string]string `mapstructure:"mapping"`
	// the error format
	Encoding string `mapstructure:"error"`
	// the response to process is a collection
	IsCollection bool `mapstructure:"is_collection"`
	// name of the field to extract to the root. If empty, the formater will do nothing
	Target string `mapstructure:"target"`
	// name of the stackMicroservices discoveryService driver to use
	DiscoveryService string `mapstructure:"discovery_service_name"`

	// list of keys to be replaced in the url
	UrlKeys []string `mapstructure:"urlKeys"`
	// number of concurrent calls this endpoint must send to the api
	ConcurrentCalls int `mapstructure:"concurrentCalls"`
	// timeout of this backend
	Timeout time.Duration `mapstructure:"timeout"`
}

func NewBackendConfig(group string, method string, host []string, hostSanitizationDisabled bool, url string, blacklist []string, whitelist []string, mapping map[string]string, encoding string, isCollection bool, target string, discoveryService string, urlKeys []string, concurrentCalls int,
	timeout time.Duration) (*httpModelConfigAVA.Backend, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewBackendConfigDefault() (*httpModelConfigAVA.Backend, *errorAVA.Error) {
	panic("Not implemented.")
}

func (b *BackendConfig) ReadLocal(fileName string) (*httpModelConfigAVA.Backend, *errorAVA.Error) {
	panic("Not implemented.")
}

func (b *BackendConfig) Parser(disableStrictREST bool, urlPattern string) (*httpModelConfigAVA.Backend, *errorAVA.Error) {
	if err := b.checkURLPattern(); err != nil {
		return nil, err
	}

	if err := b.checkParams(disableStrictREST, urlPattern); err != nil {
		return nil, err
	}

	method, err := httpModelConfigAVA.ParseHTTPVerbType(b.Method)
	if err != nil {
		return nil, errorConfigAVA.HTTPVerbWrong(err, fmt.Sprintf("Method: %s", b.Method))
	}

	encoding, errEncoder := b.checkEncode(b.Encoding)
	if errEncoder != nil {
		return nil, errEncoder
	}

	if b.ConcurrentCalls == 0 {
		b.ConcurrentCalls = ConcurrentCallsDefault
	}

	if b.Timeout == 0 {
		b.Timeout = DefaultTimeout
	}

	return httpModelConfigAVA.NewBackend(
		b.Group,
		method,
		b.Host,
		b.HostSanitizationDisabled,
		b.Url,
		b.Blacklist,
		b.Whitelist,
		b.Mapping,
		encoding,
		b.IsCollection,
		b.Target,
		b.DiscoveryService,
		b.UrlKeys,
		b.ConcurrentCalls,
		b.Timeout,
	)
}

func (b *BackendConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(b)
}

func (b *BackendConfig) checkEncode(encoder string) (httpModelConfigAVA.OutputEncodingType, *errorAVA.Error) {
	var encoding httpModelConfigAVA.OutputEncodingType
	if encoder == "" {
		encoding = httpModelConfigAVA.OutputEncodingTypeJSON
	} else {
		var err error
		encoding, err = httpModelConfigAVA.ParseOutputEncodingType(encoder)
		if err != nil {
			return 0, errorConfigAVA.OutputEncodingWrong(err, fmt.Sprintf("encoding: %s", encoder))
		}
	}
	return encoding, nil
}

func (b *BackendConfig) checkURLPattern() *errorAVA.Error {
	if b.Url == "" {
		return errorConfigAVA.URLIsEmpty(nil, fmt.Sprintf("%v", b))
	} else {
		b.Url = uriUtilsAVA.CleanPath(b.Url)
	}

	return nil
}

func (b *BackendConfig) extractPlaceHoldersFromURLTemplate(subject string, pattern *regexp.Regexp) []string {
	matches := pattern.FindAllStringSubmatch(subject, -1)
	keys := make([]string, len(matches))
	for k, v := range matches {
		keys[k] = v[1]
	}
	return keys
}

func (b *BackendConfig) paramExtractionPattern(disableStrictREST bool) *regexp.Regexp {
	if disableStrictREST {
		return simpleURLKeysPattern
	}
	return uriUtilsAVA.EndpointURLKeysPattern
}

func (b *BackendConfig) checkParams(disableStrictREST bool, urlPattern string) *errorAVA.Error {
	inputParams := b.extractPlaceHoldersFromURLTemplate(urlPattern, b.paramExtractionPattern(disableStrictREST))
	inputSet := map[string]interface{}{}
	for ip := range inputParams {
		inputSet[inputParams[ip]] = nil
	}

	b.Url = uriUtilsAVA.CleanPath(b.Url)

	outputParams := b.extractPlaceHoldersFromURLTemplate(b.Url, simpleURLKeysPattern)

	outputSet := map[string]interface{}{}
	for op := range outputParams {
		outputSet[outputParams[op]] = nil
	}

	if len(outputSet) > len(inputParams) {
		return errorGeneralAVA.UnknownError(nil, fmt.Sprintf("Too many output params! input: %v, output: %v\n", outputSet, outputParams))
	}

	tmp := b.Url
	b.UrlKeys = make([]string, len(outputParams))
	for o := range outputParams {
		if _, ok := inputSet[outputParams[o]]; !ok {
			return errorGeneralAVA.UnknownError(nil, fmt.Sprintf("Undefined output param [%s]! input: %v, output: %v\n", outputParams[o], inputParams, outputParams))
		}
		tmp = strings.Replace(tmp, "{"+outputParams[o]+"}", "{{."+strings.Title(outputParams[o])+"}}", -1)
		b.UrlKeys = append(b.UrlKeys, strings.Title(outputParams[o]))
	}
	b.Url = tmp

	return nil
}
