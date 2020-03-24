package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// InvalidConfig is a AVA Error
func InvalidConfig(e error, details interface{}) *errorAVA.Error {
	return InvalidConfigSkip(e, details, errorAVA.RetrieveCallDefault)
}

// InvalidConfigSkip is a AVA Error
func InvalidConfigSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupLogger,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusInvalidConfig,
		Message:  statusTextFunc(statusInvalidConfig),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
