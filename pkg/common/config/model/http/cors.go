package http

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type CORS struct {
	allowOrigins     []string
	exposeHeaders    []HTTPHeaderType
	maxAge           int
	allowMethods     []HTTPVerbType
	allowHeaders     []HTTPHeaderType
	allowCredentials bool
	debug            bool
}

func NewCORS(allowOrigins []string, exposeHeaders []HTTPHeaderType, maxAge int, allowMethods []HTTPVerbType, allowHeaders []HTTPHeaderType, allowCredentials bool, debug bool) (*CORS, *errorAVA.Error) {
	return &CORS{
		allowOrigins:     allowOrigins,
		exposeHeaders:    exposeHeaders,
		maxAge:           maxAge,
		allowMethods:     allowMethods,
		allowHeaders:     allowHeaders,
		allowCredentials: allowCredentials,
		debug:            debug,
	}, nil
}

func NewCORSDefault() (*CORS, *errorAVA.Error) {
	panic("Not implemented.")
}

func (cors *CORS) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(cors)
}
