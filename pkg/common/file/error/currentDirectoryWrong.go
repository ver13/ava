package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CurrentDirectoryPathWrong is a AVA Error
func CurrentDirectoryPathWrong(e error, details interface{}) *errorAVA.Error {
	return CurrentDirectoryPathWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CurrentDirectoryPathWrongSkip is a AVA Error
func CurrentDirectoryPathWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusCurrentDirectoryPathWrong,
		Message:  statusTextFunc(statusCurrentDirectoryPathWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
