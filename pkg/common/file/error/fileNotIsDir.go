package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FileNotIsDir is a AVA Error
func FileNotIsDir(e error, details interface{}) *errorAVA.Error {
	return FileNotIsDirSkip(e, details, errorAVA.RetrieveCallDefault)
}

// FileNotIsDirSkip is a AVA Error
func FileNotIsDirSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusFileNotIsDir,
		Message:  statusTextFunc(statusFileNotIsDir),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
