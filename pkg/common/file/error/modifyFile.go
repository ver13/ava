package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ModifyFile is a AVA Error
func ModifyFile(e error, details interface{}) *errorAVA.Error {
	return ModifyFileSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ModifyFileSkip is a AVA Error
func ModifyFileSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupFile,
		Code:     statusModifyFile,
		Message:  statusTextFunc(statusModifyFile),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
