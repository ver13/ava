package http

import (
	"fmt"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type CORSConfig struct {
	allowOrigins     []string `mapstructure:"allow_origins,omitempty"`
	exposeHeaders    []string `mapstructure:"expose_headers,omitempty"`
	maxAge           int      `mapstructure:"max_age,omitempty"`
	allowMethods     []string `mapstructure:"allow_methods,omitempty"`
	allowHeaders     []string `mapstructure:"allow_headers,omitempty"`
	allowCredentials bool     `mapstructure:"allow_credentials,omitempty"`
	debug            bool     `mapstructure:"debug,omitempty"`
}

func NewCORSConfig(allowOrigins []string, exposeHeaders []string, maxAge int, allowMethods []string, allowHeaders []string, allowCredentials bool, debug bool) (httpModelConfigAVA.CORS, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewCORSConfigDefault() (httpModelConfigAVA.CORS, *errorAVA.Error) {
	panic("Not implemented.")
}

func (cors *CORSConfig) ReadLocal(fileName string) (*httpModelConfigAVA.CORS, *errorAVA.Error) {
	panic("Not implemented.")
}

func (cors *CORSConfig) Parser() (*httpModelConfigAVA.CORS, *errorAVA.Error) {
	var maxAge = cors.maxAge
	// Maximum of 10 minutes.
	if cors.maxAge > 600 {
		maxAge = 600
	}

	allowHeaders, errAllowHeaders := cors.parseAllowedHeaders()
	if errAllowHeaders != nil {
		return nil, errAllowHeaders
	}

	exposeHeaders, errExposeHeaders := cors.parseExposeHeaders()
	if errExposeHeaders != nil {
		return nil, errExposeHeaders
	}

	allowMethods, errAllowMethods := cors.parseAllowedMethods()
	if errAllowMethods != nil {
		return nil, errAllowMethods
	}

	if len(cors.allowOrigins) == 0 {
		cors.allowOrigins = allowedOriginsDefault
	}

	return httpModelConfigAVA.NewCORS(
		cors.allowOrigins,
		exposeHeaders,
		maxAge,
		allowMethods,
		allowHeaders,
		cors.allowCredentials,
		cors.debug,
	)
}

func (cors *CORSConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(cors)
}

var (
	allowedHeadersDefault = []string{"Origin", "Accept", "Content-Type"}
	allowedMethodsDefault = []string{"GET", "POST"}
	allowedOriginsDefault = []string{"*"}
)

func (cors *CORSConfig) parseAllowedMethods() ([]httpModelConfigAVA.HTTPVerbType, *errorAVA.Error) {
	l := len(cors.allowMethods)
	if l == 0 {
		cors.allowMethods = allowedMethodsDefault
	}

	allowMethods := make([]httpModelConfigAVA.HTTPVerbType, len(cors.allowMethods))
	for i := 0; i < len(cors.allowMethods); i++ {
		allowMethod, err := httpModelConfigAVA.ParseHTTPVerbType(cors.allowMethods[i])
		if err != nil {
			return nil, errorConfigAVA.AllowMethodWrong(err, fmt.Sprintf("Method: %s", cors.allowMethods[i]))
		} else {
			allowMethods[i] = allowMethod
		}
	}
	return allowMethods, nil
}

func (cors *CORSConfig) parseAllowedHeaders() ([]httpModelConfigAVA.HTTPHeaderType, *errorAVA.Error) {
	l := len(cors.allowHeaders)
	if l == 0 {
		cors.allowHeaders = allowedHeadersDefault
	}

	allowHeaders := make([]httpModelConfigAVA.HTTPHeaderType, len(cors.allowHeaders))
	for i := 0; i < len(cors.allowHeaders); i++ {
		allowHeader, err := httpModelConfigAVA.ParseHTTPHeaderType(cors.allowHeaders[i])
		if err != nil {
			return nil, err
		} else {
			allowHeaders[i] = allowHeader
		}
	}
	return allowHeaders, nil
}

func (cors *CORSConfig) parseExposeHeaders() ([]httpModelConfigAVA.HTTPHeaderType, *errorAVA.Error) {
	exposeHeaders := make([]httpModelConfigAVA.HTTPHeaderType, len(cors.exposeHeaders))
	for i := 0; i < len(cors.exposeHeaders); i++ {
		exposeHeader, err := httpModelConfigAVA.ParseHTTPHeaderType(cors.exposeHeaders[i])
		if err != nil {
			return nil, err
		} else {
			exposeHeaders[i] = exposeHeader
		}
	}
	return exposeHeaders, nil
}
