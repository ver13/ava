package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// TLSVersionWrong is a AVA Error
func TLSVersionWrong(e error, details interface{}) *errorAVA.Error {
	return TLSVersionWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// TLSVersionWrongSkip is a AVA Error
func TLSVersionWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusTLSVersionWrong,
		Message:  statusTextFunc(statusTLSVersionWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
