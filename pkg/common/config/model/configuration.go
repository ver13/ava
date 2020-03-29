package model

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	versionAVA "github.com/ver13/ava/pkg/common/version"
)

type Configuration struct {
	ProjectName string
	Author      string
	Copyright   string

	Version *versionAVA.VersionInfo

	Environment *Environment
}

func (c *Configuration) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	s, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return s.Serializer(c)
}
