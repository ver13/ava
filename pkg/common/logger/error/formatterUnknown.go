package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FormatterUnknown is a AVA Error
func FormatterUnknown(e error, details interface{}) *errorAVA.Error {
	return FormatterUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// FormatterUnknownSkip is a AVA Error
func FormatterUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusLoggerFormatterUnknown,
		Message:  statusTextFunc(statusLoggerFormatterUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}

