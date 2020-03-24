package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// GenerateRandomSerialNumber is a AVA Error
func GenerateRandomSerialNumber(e error, details interface{}) *errorAVA.Error {
	return GenerateRandomSerialNumberSkip(e, details, errorAVA.RetrieveCallDefault)
}

// GenerateRandomSerialNumberSkip is a AVA Error
func GenerateRandomSerialNumberSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusGenerateRandomSerialNumberWrong,
		Message:  statusTextFunc(statusGenerateRandomSerialNumberWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
