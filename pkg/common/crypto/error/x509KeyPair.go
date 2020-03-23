package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// X509KeyPair is a AVA Error
func X509KeyPair(e error, details interface{}) *errorAVA.Error {
	return X509KeyPairSkip(e, details, errorAVA.RetrieveCallDefault)
}

// X509KeyPairSkip is a AVA Error
func X509KeyPairSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusX509KeyPair,
		Message:  statusTextFunc(statusX509KeyPair),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
