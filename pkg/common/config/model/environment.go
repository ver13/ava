package model

import (
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Environment struct {
	environmentType EnvironmentType
	name            string
	logger          *loggerAVA.Logger
	api             *httpModelConfigAVA.API
	tls             *httpModelConfigAVA.TLS
}

func NewEnvironment(environmentType EnvironmentType, name string, logger *loggerAVA.Logger, api *httpModelConfigAVA.API, tls *httpModelConfigAVA.TLS) (*Environment, *errorAVA.Error) {
	return &Environment{
		environmentType: environmentType,
		name:            name,
		logger:          logger,
		api:             api,
		tls:             tls,
	}, nil
}

func (environment *Environment) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(environment)
}
