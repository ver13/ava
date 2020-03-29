package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// AllowMethodWrong is a AVA Error
func AllowMethodWrong(e error, details interface{}) *errorAVA.Error {
	return AllowMethodWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// AllowMethodWrongSkip is a AVA Error
func AllowMethodWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusAllowMethodWrong,
		Message:  statusTextFunc(statusAllowMethodWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
