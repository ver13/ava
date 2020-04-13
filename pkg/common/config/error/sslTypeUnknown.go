package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// DialectTypeUnknown is a AVA Error
func DialectTypeUnknown(e error, details interface{}) *errorAVA.Error {
	return DialectTypeUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// DialectTypeUnknownSkip is a AVA Error
func DialectTypeUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusDialectTypeUnknown,
		Message:  statusTextFunc(statusDialectTypeUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
