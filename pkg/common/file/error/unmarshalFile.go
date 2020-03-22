package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// UnmarshalFile is a AVA Error
func UnmarshalFile(e error, details interface{}) *errorAVA.Error {
	return UnmarshalFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// UnmarshalFileSkip is a AVA Error
func UnmarshalFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusUnmarshalFile,
		Message:  statusTextFunc(statusUnmarshalFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
