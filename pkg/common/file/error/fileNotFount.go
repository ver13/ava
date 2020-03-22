package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FileNotFount is a AVA Error
func FileNotFount(e error, details interface{}) *errorAVA.Error {
	return FileNotFountSkip(e, details, errorAVA.RetrieveCallDefault)
}

// FileNotFountSkip is a AVA Error
func FileNotFountSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusFileNotFount,
		Message:  statusTextFunc(statusFileNotFount),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
