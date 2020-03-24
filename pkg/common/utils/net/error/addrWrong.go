package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// AddrWrong is a AVA Error
func AddrWrong(e error, details interface{}) *errorAVA.Error {
	return AddrWrongSkip(e, details, 3)
}

// AddrWrongSkip is a AVA Error
func AddrWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUtils,
		Subgroup: errorAVA.SubgroupNET,
		Code:     statusAddrWrong,
		Message:  statusTextFunc(statusAddrWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
