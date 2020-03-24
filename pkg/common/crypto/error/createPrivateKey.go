package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CreatePrivateKey is a AVA Error
func CreatePrivateKey(e error, details interface{}) *errorAVA.Error {
	return CreatePrivateKeySkip(e, details, errorAVA.RetrieveCallDefault)
}

// CreatePrivateKeySkip is a AVA Error
func CreatePrivateKeySkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusCreatePrivateKey,
		Message:  statusTextFunc(statusCreatePrivateKey),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
