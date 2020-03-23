package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// PasswordIsEmpty is a AVA Error
func PasswordIsEmpty(e error, details interface{}) *errorAVA.Error {
	return PasswordIsEmptySkip(e, details, errorAVA.RetrieveCallDefault)
}

// PasswordIsEmptySkip is a AVA Error
func PasswordIsEmptySkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusPasswordIsEmpty,
		Message:  statusTextFunc(statusPasswordIsEmpty),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
