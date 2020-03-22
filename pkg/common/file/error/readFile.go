package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ReadFile is a AVA Error
func ReadFile(e error, details interface{}) *errorAVA.Error {
	return ReadFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ReadFileSkip is a AVA Error
func ReadFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusReadFile,
		Message:  statusTextFunc(statusReadFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
