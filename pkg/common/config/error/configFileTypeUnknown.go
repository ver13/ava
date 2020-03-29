package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ConfigFileTypeUnknown is a AVA Error
func ConfigFileTypeUnknown(e error, details interface{}) *errorAVA.Error {
	return ConfigFileTypeUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ConfigFileTypeUnknownSkip is a AVA Error
func ConfigFileTypeUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusConfigFileTypeUnknown,
		Message:  statusTextFunc(statusConfigFileTypeUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
