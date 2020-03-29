package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// HTTPHeaderUnknown is a AVA Error
func HTTPHeaderUnknown(e error, details interface{}) *errorAVA.Error {
	return HTTPHeaderUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// HTTPHeaderUnknownSkip is a AVA Error
func HTTPHeaderUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusHTTPHeaderUnknown,
		Message:  statusTextFunc(statusHTTPHeaderUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
