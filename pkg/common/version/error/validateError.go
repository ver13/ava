package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func ValidateError(e error, details interface{}) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupVersion,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusValidateError,
		Message:  statusTextFunc(statusValidateError),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}