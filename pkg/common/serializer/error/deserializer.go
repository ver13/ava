package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Deserializer is a AVA Error
func Deserializer(e error, details interface{}) *errorAVA.Error {
	return DeserializerSkip(e, details, errorAVA.RetrieveCallDefault)
}

// DeserializerSkip is a AVA Error
func DeserializerSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupSerializer,
		Code:     statusDeserializer,
		Message:  statusTextFunc(statusDeserializer),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
