package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func TypeUnknown(e error, details interface{}) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupHash,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusTypeUnknown,
		Message:  statusTextFunc(statusTypeUnknown),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
