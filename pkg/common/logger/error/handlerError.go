package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// HandlerError is a AVA Error
func HandlerError(e error, details interface{}) *errorAVA.Error {
	return HandlerErrorSkip(e, details, errorAVA.RetrieveCallDefault)
}

// HandlerErrorSkip is a AVA Error
func HandlerErrorSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusHandlerError,
		Message:  statusTextFunc(statusHandlerError),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
