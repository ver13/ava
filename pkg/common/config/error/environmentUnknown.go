package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func EnvironmentUnknown(e error, details interface{}) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusEnvironmentUnknown,
		Message:  statusTextFunc(statusEnvironmentUnknown),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
