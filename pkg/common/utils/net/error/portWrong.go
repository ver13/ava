package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// PortWrong is a AVA Error
func PortWrong(e error, details interface{}) *errorAVA.Error {
	return PortWrongSkip(e, details, 3)
}

// PortWrongSkip is a AVA Error
func PortWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUtils,
		Subgroup: errorAVA.SubgroupNET,
		Code:     statusPortWrong,
		Message:  statusTextFunc(statusPortWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
