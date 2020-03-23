package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CreateCertificate is a AVA Error
func CreateCertificate(e error, details interface{}) *errorAVA.Error {
	return CreateCertificateSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CreateCertificateSkip is a AVA Error
func CreateCertificateSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusCreateCertificate,
		Message:  statusTextFunc(statusCreateCertificate),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
