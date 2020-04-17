package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// WritingDataFailed is a AVA Error
func WritingDataFailed(e error, details interface{}) *errorAVA.Error {
	return WritingDataFailedSkip(e, details, errorAVA.RetrieveCallDefault)
}

// WritingDataFailedSkip is a AVA Error
func WritingDataFailedSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupTools,
		Subgroup: errorAVA.SubgroupAVAEnum,
		Code:     statusWritingDataFailed,
		Message:  statusTextFunc(statusWritingDataFailed),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
