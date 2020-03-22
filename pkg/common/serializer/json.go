package serializer

import (
	"encoding/json"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

type JSON struct {
}

func (j *JSON) Serializer(data interface{}) ([]byte, *errorAVA.Error) {
	var result []byte
	var err error

	result, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, errorSerializerAVA.Serializer(err, "Failed to generate JSON.")
	}

	return result, nil
}

func (j *JSON) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	if err := json.Unmarshal(data, out); err != nil {
		return errorSerializerAVA.Deserializer(err, "Failed to generate struct.")
	}
	return nil
}

func (j *JSON) String() string {
	return j.Type().String()
}

func (j *JSON) Type() SerializerType {
	return SerializerTypeJson
}
