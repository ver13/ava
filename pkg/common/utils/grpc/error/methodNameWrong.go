package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// StatusMethodNameWrong is a AVA Error
func StatusMethodNameWrong(e error, details interface{}) *errorAVA.Error {
	return StatusMethodNameWrongSkip(e, details, 3)
}

// StatusMethodNameWrongSkip is a AVA Error
func StatusMethodNameWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUtils,
		Subgroup: errorAVA.SubgroupGRPC,
		Code:     statusMethodNameWrong,
		Message:  statusTextFunc(statusMethodNameWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
