package serializer

import (
	"encoding/xml"
	"sync"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

var 	onceXML    sync.Once

type XML struct {
}

func init() {
	onceXML.Do(func() {
		Register(SerializerTypeXml, &XML{})
	})
}

func (x *XML) Serializer(data interface{}) ([]byte, *errorAVA.Error) {
	var result []byte
	var err error

	result, err = xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, errorSerializerAVA.Serializer(err, "Failed to generate XML.")
	}

	return result, nil
}

func (x *XML) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	if err := xml.Unmarshal(data, out); err != nil {
		return errorSerializerAVA.Deserializer(err, "Failed to generate struct.")
	}
	return nil
}

func (x *XML) String() string {
	return x.Type().String()
}

func (x *XML) Type() SerializerType {
	return SerializerTypeXml
}
