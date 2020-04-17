package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// NoDocOnEnum is a AVA Error
func NoDocOnEnum(e error, details interface{}) *errorAVA.Error {
	return NoDocOnEnumSkip(e, details, errorAVA.RetrieveCallDefault)
}

// NoDocOnEnumSkip is a AVA Error
func NoDocOnEnumSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupTools,
		Subgroup: errorAVA.SubgroupAVAEnum,
		Code:     statusNoDocOnEnum,
		Message:  statusTextFunc(statusNoDocOnEnum),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
