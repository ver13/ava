package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// OpenFile is a AVA Error
func OpenFile(e error, details interface{}) *errorAVA.Error {
	return OpenFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// OpenFileSkip is a AVA Error
func OpenFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusOpenFile,
		Message:  statusTextFunc(statusOpenFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
