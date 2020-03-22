package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// OpenDir is a AVA Error
func OpenDir(e error, details interface{}) *errorAVA.Error {
	return OpenDirSkip(e, details, errorAVA.RetrieveCallDefault)
}

// OpenDirSkip is a AVA Error
func OpenDirSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusOpenDir,
		Message:  statusTextFunc(statusOpenDir),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
