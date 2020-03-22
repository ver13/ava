package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Unknown is a AVA Error
func Unknown(e error, details interface{}) *errorAVA.Error {
	return UnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// UnknownSkip is a AVA Error
func UnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusUnknown,
		Message:  statusTextFunc(statusUnknown),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
