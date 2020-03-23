package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// LogLevelUnknown is a AVA Error
func LogLevelUnknown(e error, details interface{}) *errorAVA.Error {
	return LogLevelUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// LogLevelUnknownSkip is a AVA Error
func LogLevelUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusLoggerLevelUnknown,
		Message:  statusTextFunc(statusLoggerLevelUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}

