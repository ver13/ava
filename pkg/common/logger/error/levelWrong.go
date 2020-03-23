package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// InvalidConfig is a AVA Error
func LevelWrong(e error, details interface{}) *errorAVA.Error {
	return LevelWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// LevelWrongSkip is a AVA Error
func LevelWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusLevelWrong,
		Message:  statusTextFunc(statusLevelWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
