package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CloseFile is a AVA Error
func CloseFile(e error, details interface{}) *errorAVA.Error {
	return CloseFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CloseFileSkip is a AVA Error
func CloseFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusCloseFile,
		Message:  statusTextFunc(statusCloseFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
