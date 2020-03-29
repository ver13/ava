package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ConfigVersionWrong is a AVA Error
func ConfigVersionWrong(e error, details interface{}) *errorAVA.Error {
	return ConfigVersionWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ConfigVersionWrongSkip is a AVA Error
func ConfigVersionWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusConfigVersionWrong,
		Message:  statusTextFunc(statusConfigVersionWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
