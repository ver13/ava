package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func WriteBinary(e error, details string) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupHash,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusWriteBinary,
		Message:  statusTextFunc(statusWriteBinary),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
