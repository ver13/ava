package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// NotAvailable is a AVA Error
func NotAvailable(e error, details interface{}) *errorAVA.Error {
	return NotAvailableSkip(e, details, errorAVA.RetrieveCallDefault)
}

// NotAvailableSkip is a AVA Error
func NotAvailableSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusNotAvailable,
		Message:  statusTextFunc(statusNotAvailable),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}

