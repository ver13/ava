package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Timeout is a AVA Error
func Timeout(e error, details interface{}) *errorAVA.Error {
	return TimeoutSkip(e, details, 3)
}

// TimeoutSkip is a AVA Error
func TimeoutSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusTimeout,
		Message:  statusTextFunc(statusTimeout),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
