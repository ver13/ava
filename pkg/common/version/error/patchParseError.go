package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func PatchParseError(e error, details interface{}) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupVersion,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusPatchParseError,
		Message:  statusTextFunc(statusPatchParseError),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
