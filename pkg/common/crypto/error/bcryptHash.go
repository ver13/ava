package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// BcryptHash is a AVA Error
func BcryptHash(e error, details interface{}) *errorAVA.Error {
	return BcryptHashSkip(e, details, errorAVA.RetrieveCallDefault)
}

// BcryptHashSkip is a AVA Error
func BcryptHashSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusBcryptHash,
		Message:  statusTextFunc(statusBcryptHash),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
