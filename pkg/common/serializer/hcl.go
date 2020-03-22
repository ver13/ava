package serializer

import (
	"encoding/json"
	"sync"
	
	"github.com/hashicorp/hcl"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

var onceHCL    sync.Once

type HCL struct {

}

func init() {
	onceHCL.Do(func() {
		Register(SerializerTypeHcl, &HCL{})
	})
}

func (h *HCL) Serializer(data interface{}) ([]byte, *errorAVA.Error) {
	var result []byte
	var err error

	result, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, errorSerializerAVA.Serializer(err, "Failed to generate HCL.")
	}

	return result, nil
}

func (h *HCL) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	if err := hcl.Unmarshal(data, out); err != nil {
		return errorSerializerAVA.Deserializer(err, "Failed to generate struct.")
	}
	return nil
}

func (h *HCL) String() string {
	return h.Type().String()
}

func (h *HCL) Type() SerializerType {
	return SerializerTypeHcl
}
