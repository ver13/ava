package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// EnvironmentWrong is a AVA Error
func EnvironmentWrong(e error, details interface{}) *errorAVA.Error {
	return EnvironmentWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// EnvironmentWrongSkip is a AVA Error
func EnvironmentWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusEnvironmentWrong,
		Message:  statusTextFunc(statusEnvironmentWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
