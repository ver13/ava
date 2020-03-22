package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CopyDir is a AVA Error
func CopyDir(e error, details interface{}) *errorAVA.Error {
	return CopyDirSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CopyDirSkip is a AVA Error
func CopyDirSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusCopyDir,
		Message:  statusTextFunc(statusCopyDir),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
