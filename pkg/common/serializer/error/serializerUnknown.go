package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// SerializerUnknown is a AVA Error
func SerializerUnknown(e error, details interface{}) *errorAVA.Error {
	return SerializerUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// SerializerUnknownSkip is a AVA Error
func SerializerUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupSerializer,
		Code:     statusSerializerUnknown,
		Message:  statusTextFunc(statusSerializerUnknown),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
