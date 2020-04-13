package file

import (
	"fmt"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	"github.com/ver13/ava/pkg/common/config/model"
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	"github.com/ver13/ava/pkg/common/config/source/file/blockchain"
	"github.com/ver13/ava/pkg/common/config/source/file/http"
	"github.com/ver13/ava/pkg/common/config/source/file/stored"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type EnvironmentConfig struct {
	Name string `mapstructure:"name,omitempty"`
	Type string `mapstructure:"type,omitempty"`

	Logger *LoggerConfig `mapstructure:"logger,omitempty"`

	Tls *http.TLSConfig `mapstructure:"crypto,omitempty"`
	Api *http.APIConfig `mapstructure:"api,omitempty"`

	Database *stored.DbSQLConfig `mapstructure:"database,omitempty"`

	Blockchain *blockchain.Blockchain `mapstructure:"blockchain,omitempty"`

	//	Stack []*stackMicroservices.StackMicroservices `mapstructure:"stack,omitempty"`
}

func (e *EnvironmentConfig) Parser() (*model.Environment, *errorAVA.Error) {
	_type, err := model.ParseEnvironmentType(e.Type)
	if err != nil {
		return nil, errorConfigAVA.EnvironmentWrong(err.Error(), fmt.Sprintf("Environment incorrect in configuration file. Environment: %s", e.Type))
	}

	var log *loggerAVA.Logger
	if e.Logger == nil {
		log = loggerAVA.GetInstance()
	} else {
		var errLogger *errorAVA.Error
		log, errLogger = e.Logger.Parser()
		if errLogger != nil {
			log = loggerAVA.GetInstance()
		}
	}

	var tls *httpModelConfigAVA.TLS
	var errTLS *errorAVA.Error
	if e.Tls == nil {
		tls, errTLS = httpModelConfigAVA.NewTLSDefault()
	} else {
		tls, errTLS = e.Tls.Parser()
	}
	if errTLS != nil {
		return nil, errTLS
	}

	var api *httpModelConfigAVA.API
	if e.Api == nil {
		return nil, errorConfigAVA.APIIsEmpty(nil, fmt.Sprintf("%v", e))
	} else {
		var errAPI *errorAVA.Error
		api, errAPI = e.Api.Parser()
		if errAPI != nil {
			return nil, errAPI
		}
	}
	/*
		var service []*service.Microservice
		if len(environment.Stack) == 0 {
			loggerAVA.Warn("Microservices configuration dervice is empty.")
		} else {
			for _, elem := range environment.Stack {
				microservice, errParse := elem.Parser()
				if errParse != nil {
					return nil, errParse
				}
				service = append(service, microservice)
			}
		}
	*/
	return model.NewEnvironment(_type, e.Name, log, api, tls)
}

func (e *EnvironmentConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	s, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return s.Serializer(e)
}
