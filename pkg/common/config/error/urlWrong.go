package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// URLWrong is a AVA Error
func URLWrong(e error, details interface{}) *errorAVA.Error {
	return URLWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// URLWrongSkip is a AVA Error
func URLWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusURLWrong,
		Message:  statusTextFunc(statusURLWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
