package serializer

import (
	"sync"

	"gopkg.in/yaml.v2"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

var onceYAML sync.Once

type YAML struct {
}

func init() {
	onceYAML.Do(func() {
		Register(SerializerTypeYaml, (*Serializer)(&YAML{}))
	})
}

func (y *YAML) Serializer(data interface{}) ([]byte, *errorAVA.Error) {
	result, err := yaml.Marshal(data)
	if err != nil {
		return nil, errorSerializerAVA.Serializer(err, "Failed to generate YAML.")
	}

	return result, nil
}

func (y *YAML) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	if err := yaml.Unmarshal(data, out); err != nil {
		return errorSerializerAVA.Deserializer(err, "Failed to generate struct.")
	}
	return nil
}

func (y *YAML) Type() SerializerType {
	return SerializerTypeYaml
}

func (y *YAML) String() string {
	return y.Type().String()
}
