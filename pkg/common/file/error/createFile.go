package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CreateFile is a AVA Error
func CreateFile(e error, details interface{}) *errorAVA.Error {
	return CreateFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CreateFileSkip is a AVA Error
func CreateFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusCreateFile,
		Message:  statusTextFunc(statusCreateFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
