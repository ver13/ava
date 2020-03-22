package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// EncodeQRError is a AVA Error
func EncodeQRError(e error, details interface{}) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupQR,
		Err:      e,
		Details:  fmt.Sprintf("%v.", details),
		Code:     statusEncodeQRError,
		Message:  statusTextFunc(statusEncodeQRError),
		Info:     errorAVA.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
