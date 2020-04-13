package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// SSLTypeUnknown is a AVA Error
func SSLTypeUnknown(e error, details interface{}) *errorAVA.Error {
	return SSLTypeUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// SSLTypeUnknownSkip is a AVA Error
func SSLTypeUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusSSLTypeUnknown,
		Message:  statusTextFunc(statusSSLTypeUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
