package general

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// NotEqual is a AVA Error
func NotEqual(e error, details interface{}) *errorAVA.Error {
	return NotEqualSkip(e, details, errorAVA.RetrieveCallDefault)
}

// NotEqualSkip is a AVA Error
func NotEqualSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUnknown,
		Subgroup: errorAVA.SubgroupUnknown,
		Code:     NotEqualCode,
		Message:  StatusTextFunc(NotEqualCode),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
