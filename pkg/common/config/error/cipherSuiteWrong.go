package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CipherSuiteWrong is a AVA Error
func CipherSuiteWrong(e error, details interface{}) *errorAVA.Error {
	return CipherSuiteWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CipherSuiteWrongSkip is a AVA Error
func CipherSuiteWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusCipherSuiteWrong,
		Message:  statusTextFunc(statusCipherSuiteWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
