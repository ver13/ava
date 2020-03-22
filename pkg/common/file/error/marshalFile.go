package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// MarshalFile is a AVA Error
func MarshalFile(e error, details interface{}) *errorAVA.Error {
	return MarshalFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// MarshalFileSkip is a AVA Error
func MarshalFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusMarshalFile,
		Message:  statusTextFunc(statusMarshalFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
