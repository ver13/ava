package serializer

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type Serializer struct {

}

func (h *Serializer) Serializer(data interface{}) ([]byte, *errorAVA.Error) {
	panic("Not implemented.")
}

func (h *Serializer) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	panic("Not implemented.")
}

func (h *Serializer) String() string {
	panic("Not implemented.")
}

func (h *Serializer) Type() SerializerType {
	panic("Not implemented.")
}
