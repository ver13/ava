package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// NotImplemented is a AVA Error
func NotImplemented(e error, details interface{}) *errorAVA.Error {
	return NotImplementedSkip(e, details, errorAVA.RetrieveCallDefault)
}

// NotImplementedSkip is a AVA Error
func NotImplementedSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusNotImplemented,
		Message:  statusTextFunc(statusNotImplemented),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
