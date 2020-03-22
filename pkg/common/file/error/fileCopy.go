package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FileCopy is a AVA Error
func FileCopy(e error, details interface{}) *errorAVA.Error {
	return FileCopySkip(e, details, errorAVA.RetrieveCallDefault)
}

// FileCopySkip is a AVA Error
func FileCopySkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusFileCopy,
		Message:  statusTextFunc(statusFileCopy),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
