package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func BuildVersionParseError(e error, details interface{}) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupVersion,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusBuildVersionParseError,
		Message:  statusTextFunc(statusBuildVersionParseError),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
