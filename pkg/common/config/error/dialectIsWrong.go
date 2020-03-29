package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// DialectIsWrong is a AVA Error
func DialectIsWrong(e error, details interface{}) *errorAVA.Error {
	return DialectIsWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// DialectIsWrongSkip is a AVA Error
func DialectIsWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusDialectIsWrong,
		Message:  statusTextFunc(statusDialectIsWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
