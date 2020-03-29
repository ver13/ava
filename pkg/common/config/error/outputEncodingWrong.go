package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// OutputEncodingWrong is a AVA Error
func OutputEncodingWrong(e error, details interface{}) *errorAVA.Error {
	return OutputEncodingWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// OutputEncodingWrongSkip is a AVA Error
func OutputEncodingWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusOutputEncodingWrong,
		Message:  statusTextFunc(statusOutputEncodingWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
