package serializer

import (
	"fmt"
	"sync"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorSerializerAVA "github.com/ver13/ava/pkg/common/serializer/error"
)

type Factory map[SerializerType]*Serializer

var (
	factory Factory
	once    sync.Once
)

func init() {
	once.Do(func() {
		factory = make(map[SerializerType]*Serializer)
	})
}

func GetInstance() *Factory {
	return &factory
}

func (s Factory) Register(t SerializerType, serializer *Serializer) *errorAVA.Error {
	s[t] = serializer
	fmt.Sprintf("Serializer %v has registered.", t)
	return nil
}
func Register(t SerializerType, serializer *Serializer) *errorAVA.Error {
	return GetInstance().Register(t, serializer)
}

func (s Factory) SerializerFactory(t SerializerType) (*Serializer, *errorAVA.Error) {
	var serializer *Serializer

	serializer = s[t]
	if serializer == nil {
		fmt.Sprintf("Serializer is not found and return error.")
		return nil, errorSerializerAVA.NotImplemented(nil, fmt.Sprintf("Serializer type: %s", t.String()))
	}
	fmt.Sprintf("Serializer %s is registered and return it.", t)
	return serializer, nil
}
func GetSerializer(t SerializerType) *Serializer {
	serializer, err := GetInstance().SerializerFactory(t)
	if err != nil {
		return nil
	}
	return serializer
}
