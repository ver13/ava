package file

import (
	"fmt"

	"github.com/ver13/ava/pkg/common/config/model"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	versionAVA "github.com/ver13/ava/pkg/common/version"

	errorAVA "github.com/ver13/ava/pkg/common/error"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
)

type ConfigurationConfig struct {
	ProjectName string `mapstructure:"projectName"`
	Author      string `mapstructure:"author,omitempty"`
	Copyright   string `mapstructure:"copyright,omitempty"`

	EnvironmentActivated string `mapstructure:"environment_activated"`

	Environments []*EnvironmentConfig `mapstructure:"environments"`
}

func (c *ConfigurationConfig) Parser(environmentActivated model.EnvironmentType) (*model.Configuration, *errorAVA.Error) {
	if c.Environments == nil {
		return nil, errorConfigAVA.EnvironmentsIsEmpty(nil, fmt.Sprintf("%v", c))
	} else if len(c.Environments) == 0 {
		return nil, errorConfigAVA.EnvironmentsIsEmpty(nil, fmt.Sprintf("%v", c))
	}

	var found = false
	var environmentUsedIndex int = -1
	for i := 0; i <= 2; i++ {
		if len(c.Environments) > i {
			if c.Environments[i] != nil && c.Environments[i].Type == environmentActivated.String() {
				found = true
				environmentUsedIndex = i
				break
			}
		}
	}

	var err *errorAVA.Error = nil
	var environment *model.Environment = nil
	if !found {
		if environmentActivated == model.EnvironmentTypeUnknown {
			return nil, errorConfigAVA.EnvironmentWrong(nil, fmt.Sprintf("Environment activated: %s", environmentActivated.String()))
		} else {
			environment, err = c.Environments[environmentActivated].Parser()
		}
	} else {
		environment, err = c.Environments[environmentUsedIndex].Parser()
	}
	if err != nil {
		return nil, err
	}

	return &model.Configuration{
		ProjectName: c.ProjectName,
		Author:      c.Author,
		Copyright:   c.Copyright,
		Version:     versionAVA.GetInstance(),
		Environment: environment,
	}, nil
}

func (c *ConfigurationConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	s, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return s.Serializer(c)
}
