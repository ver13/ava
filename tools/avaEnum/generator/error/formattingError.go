package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FormattingError is a AVA Error
func FormattingError(e error, details interface{}) *errorAVA.Error {
	return FormattingErrorSkip(e, details, errorAVA.RetrieveCallDefault)
}

// FormattingErrorSkip is a AVA Error
func FormattingErrorSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupTools,
		Subgroup: errorAVA.SubgroupAVAEnum,
		Code:     statusFormattingError,
		Message:  statusTextFunc(statusFormattingError),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
