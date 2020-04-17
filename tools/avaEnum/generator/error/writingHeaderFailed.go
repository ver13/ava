package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// WritingHeaderFailed is a AVA Error
func WritingHeaderFailed(e error, details interface{}) *errorAVA.Error {
	return WritingHeaderFailedSkip(e, details, errorAVA.RetrieveCallDefault)
}

// WritingHeaderFailedSkip is a AVA Error
func WritingHeaderFailedSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupTools,
		Subgroup: errorAVA.SubgroupAVAEnum,
		Code:     statusWritingHeaderFailed,
		Message:  statusTextFunc(statusWritingHeaderFailed),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
