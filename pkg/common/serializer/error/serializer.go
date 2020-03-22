package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Serializer is a AVA Error
func Serializer(e error, details interface{}) *errorAVA.Error {
	return SerializerSkip(e, details, errorAVA.RetrieveCallDefault)
}

// SerializerSkip is a AVA Error
func SerializerSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupSerializer,
		Code:     statusSerializer,
		Message:  statusTextFunc(statusSerializer),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
