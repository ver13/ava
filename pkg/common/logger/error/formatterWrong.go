package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FormatterWrong is a AVA Error
func FormatterWrong(e error, details interface{}) *errorAVA.Error {
	return FormatterWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// FormatterWrongSkip is a AVA Error
func FormatterWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusFormatterWrong,
		Message:  statusTextFunc(statusFormatterWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
