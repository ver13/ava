package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// GenerateKeyWrong is a AVA Error
func GenerateKeyWrong(e error, details interface{}) *errorAVA.Error {
	return GenerateKeyWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// GenerateKeyWrongSkip is a AVA Error
func GenerateKeyWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupGeneral,
		Code:     statusGenerateKeyWrong,
		Message:  statusTextFunc(statusGenerateKeyWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
