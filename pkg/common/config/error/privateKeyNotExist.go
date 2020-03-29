package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// PrivateKeyNotExist is a AVA Error
func PrivateKeyNotExist(e error, details interface{}) *errorAVA.Error {
	return PrivateKeyNotExistSkip(e, details, errorAVA.RetrieveCallDefault)
}

// PrivateKeyNotExistSkip is a AVA Error
func PrivateKeyNotExistSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusPrivateKeyNotExist,
		Message:  statusTextFunc(statusPrivateKeyNotExist),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
