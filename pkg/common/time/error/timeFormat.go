package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// TimeFormat is a AVA Error
func TimeFormat(e error, details interface{}) *errorAVA.Error {
	return TimeFormatSkip(e, details, 3)
}

// TimeFormatSkip is a AVA Error
func TimeFormatSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusTimeFormat,
		Message:  statusTextFunc(statusTimeFormat),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
