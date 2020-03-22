package serializer

import (
	"bytes"
	"fmt"
	"sync"
	
	"github.com/BurntSushi/toml"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

var 	onceTOML    sync.Once

type TOML struct {
}

func init() {
	onceTOML.Do(func() {
		Register(SerializerTypeToml, &TOML{})
	})
}

func (t *TOML) Serializer(data interface{}) (out []byte, err *errorAVA.Error) {
	defer func() {
		if r := recover(); r != nil {
			err = errorSerializerAVA.Serializer(nil, fmt.Sprintf("The data could not be marshalled to TOML. %v", r))
		}
	}()

	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	errEncode := enc.Encode(data)
	if err != nil {
		return nil, errorSerializerAVA.Serializer(errEncode, "Failed to generate TOML.")
	}
	out = buf.Bytes()

	return out, nil
}

func (t *TOML) Deserializer(data []byte, out interface{}) *errorAVA.Error {
	err := toml.Unmarshal(data, out)
	if err != nil {
		return errorSerializerAVA.Deserializer(err, "Failed to generate struct.")
	}
	return nil
}

func (t *TOML) String() string {
	return t.Type().String()
}

func (t *TOML) Type() SerializerType {
	return SerializerTypeToml
}
