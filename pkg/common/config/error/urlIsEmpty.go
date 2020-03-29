package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// URLIsEmpty is a AVA Error
func URLIsEmpty(e error, details interface{}) *errorAVA.Error {
	return URLIsEmptySkip(e, details, errorAVA.RetrieveCallDefault)
}

// URLIsEmptySkip is a AVA Error
func URLIsEmptySkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusURLIsEmpty,
		Message:  statusTextFunc(statusURLIsEmpty),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
