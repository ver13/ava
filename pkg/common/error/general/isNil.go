package general

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// IsNil is a AVA Error
func IsNil(e error, details interface{}) *errorAVA.Error {
	return IsNilSkip(e, details, errorAVA.RetrieveCallDefault)
}

// IsNilSkip is a AVA Error
func IsNilSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     IsNilCode,
		Message:  StatusTextFunc(IsNilCode),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
