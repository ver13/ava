package model

import (
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Environment struct {
	Type   EnvironmentType
	Name   string
	Logger *loggerAVA.Logger
	API    *httpModelConfigAVA.API
	TLS    *httpModelConfigAVA.TLS
}

func NewEnvironment(environmentType EnvironmentType, name string, logger *loggerAVA.Logger, api *httpModelConfigAVA.API, tls *httpModelConfigAVA.TLS) (*Environment, *errorAVA.Error) {
	return &Environment{
		Type:   environmentType,
		Name:   name,
		Logger: logger,
		API:    api,
		TLS:    tls,
	}, nil
}

func (environment *Environment) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(environment)
}
