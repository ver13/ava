package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// HostPortWrong is a AVA Error
func HostPortWrong(e error, details interface{}) *errorAVA.Error {
	return HostPortWrongSkip(e, details, 3)
}

// HostPortWrongSkip is a AVA Error
func HostPortWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUtils,
		Subgroup: errorAVA.SubgroupNET,
		Code:     statusHostPortWrong,
		Message:  statusTextFunc(statusHostPortWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}

