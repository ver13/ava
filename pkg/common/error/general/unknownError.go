package general

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// UnknownError is a AVA Error
func UnknownError(e error, details interface{}) *errorAVA.Error {
	return UnknownErrorSkip(e, details, errorAVA.RetrieveCallDefault)
}

// UnknownErrorSkip is a AVA Error
func UnknownErrorSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     UnknownErrorCode,
		Message:  StatusTextFunc(UnknownErrorCode),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
