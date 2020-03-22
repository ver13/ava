package error

import (
"fmt"

errorAVA "github.com/ver13/ava/pkg/common/error"
)

// NotFoundData is a AVA Error
func NotFoundData(e error, details interface{}) *errorAVA.Error {
	return NotFoundDataSkip(e, details, 3)
}

// NotFoundDataSkip is a AVA Error
func NotFoundDataSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupValidator,
		Code:     statusNotFoundData,
		Message:  statusTextFunc(statusNotFoundData),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
