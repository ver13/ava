package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ListDir is a AVA Error
func ListDir(e error, details interface{}) *errorAVA.Error {
	return ListDirSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ListDirSkip is a AVA Error
func ListDirSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusListDir,
		Message:  statusTextFunc(statusListDir),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
