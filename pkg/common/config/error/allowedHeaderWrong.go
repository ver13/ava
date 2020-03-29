package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// AllowedHeaderWrong is a AVA Error
func AllowedHeaderWrong(e error, details interface{}) *errorAVA.Error {
	return AllowedHeaderWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// AllowedHeaderWrongSkip is a AVA Error
func AllowedHeaderWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusAllowedHeaderWrong,
		Message:  statusTextFunc(statusAllowedHeaderWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
