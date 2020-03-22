package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CreateDir is a AVA Error
func CreateDir(e error, details interface{}) *errorAVA.Error {
	return CreateDirSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CreateDirSkip is a AVA Error
func CreateDirSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusCreateDir,
		Message:  statusTextFunc(statusCreateDir),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
