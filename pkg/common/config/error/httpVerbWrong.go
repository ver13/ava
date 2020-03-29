package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// HTTPVerbWrong is a AVA Error
func HTTPVerbWrong(e error, details interface{}) *errorAVA.Error {
	return HTTPVerbWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// HTTPVerbWrongSkip is a AVA Error
func HTTPVerbWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusHTTPVerbWrong,
		Message:  statusTextFunc(statusHTTPVerbWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
