package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// PublicKeyNotExist is a AVA Error
func PublicKeyNotExist(e error, details interface{}) *errorAVA.Error {
	return PublicKeyNotExistSkip(e, details, errorAVA.RetrieveCallDefault)
}

// PublicKeyNotExistSkip is a AVA Error
func PublicKeyNotExistSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusPublicKeyNotExist,
		Message:  statusTextFunc(statusPublicKeyNotExist),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
