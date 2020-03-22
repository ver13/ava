package serializer

import (
	"fmt"
	"sync"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

type SerializerFactory map[SerializerType]SerializerI

var (
	factory SerializerFactory
	once    sync.Once
)

func init() {
	once.Do(func() {
		factory = make(map[SerializerType]SerializerI)

		Register(SerializerTypeHcl, &HCL{})
		Register(SerializerTypeJson, &JSON{})
		Register(SerializerTypeProto, &Proto{})
		Register(SerializerTypeToml, &TOML{})
		Register(SerializerTypeXml, &XML{})
		Register(SerializerTypeYaml, &YAML{})
	})
}

func GetInstance() SerializerFactory {
	return factory
}

func (s SerializerFactory) Register(t SerializerType, serializer SerializerI) *errorAVA.Error {
	s[t] = serializer
	return nil
}
func Register(t SerializerType, serializer SerializerI) *errorAVA.Error {
	return GetInstance().Register(t, serializer)
}

func (s SerializerFactory) SerializerFactory(t SerializerType) (SerializerI, *errorAVA.Error) {
	var serializer SerializerI

	serializer = s[t]
	if serializer == nil {
		return nil, errorSerializerAVA.NotImplemented(nil, fmt.Sprintf("Serializer type: %s", t.String()))
	}
	return serializer, nil
}
func GetSerializer(t SerializerType) SerializerI {
	serializer, err := GetInstance().SerializerFactory(t)
	if err != nil {
		return nil
	}
	return serializer
}
