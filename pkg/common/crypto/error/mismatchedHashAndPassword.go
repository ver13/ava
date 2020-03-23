package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// MismatchedHashAndPassword is a AVA Error
func MismatchedHashAndPassword(e error, details interface{}) *errorAVA.Error {
	return MismatchedHashAndPasswordSkip(e, details, errorAVA.RetrieveCallDefault)
}

// MismatchedHashAndPasswordSkip is a AVA Error
func MismatchedHashAndPasswordSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusMismatchedHashAndPassword,
		Message:  statusTextFunc(statusMismatchedHashAndPassword),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
