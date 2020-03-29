package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ExposedHeaderWrong is a AVA Error
func ExposedHeaderWrong(e error, details interface{}) *errorAVA.Error {
	return ExposedHeaderWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ExposedHeaderWrongSkip is a AVA Error
func ExposedHeaderWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusExposedHeaderWrong,
		Message:  statusTextFunc(statusExposedHeaderWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
