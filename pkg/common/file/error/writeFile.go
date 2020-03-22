package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// WriteFile is a AVA Error
func WriteFile(e error, details interface{}) *errorAVA.Error {
	return WriteFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// WriteFileSkip is a AVA Error
func WriteFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusWriteFile,
		Message:  statusTextFunc(statusWriteFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
