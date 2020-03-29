package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// EnvironmentsIsEmpty is a AVA Error
func EnvironmentsIsEmpty(e error, details interface{}) *errorAVA.Error {
	return EnvironmentsIsEmptySkip(e, details, errorAVA.RetrieveCallDefault)
}

// EnvironmentsIsEmptySkip is a AVA Error
func EnvironmentsIsEmptySkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusEnvironmentsIsEmpty,
		Message:  statusTextFunc(statusEnvironmentsIsEmpty),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
