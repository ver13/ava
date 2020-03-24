package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// IPNotFount is a AVA Error
func IPNotFount(e error, details interface{}) *errorAVA.Error {
	return IPNotFountSkip(e, details, 3)
}

// IPNotFountSkip is a AVA Error
func IPNotFountSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusIPNotFount,
		Message:  statusTextFunc(statusIPNotFount),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
