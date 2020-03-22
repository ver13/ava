package serializer

import (
	"github.com/golang/protobuf/proto"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

type Proto struct {
}

func (p *Proto) Serializer(data interface{}) ([]byte, *errorAVA.Error) {
	var result []byte
	var err error

	result, err = proto.Marshal(data.(proto.Message))
	if err != nil {
		return nil, errorSerializerAVA.Serializer(err, "Failed to generate Proto.")
	}

	return result, nil
}

func (p *Proto) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	if err := proto.Unmarshal(data, out.(proto.Message)); err != nil {
		return errorSerializerAVA.Deserializer(err, "Failed to generate struct.")
	}
	return nil
}

func (p *Proto) String() string {
	return p.Type().String()
}

func (p *Proto) Type() SerializerType {
	return SerializerTypeProto
}
