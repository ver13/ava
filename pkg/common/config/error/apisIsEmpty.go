package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// APIIsEmpty is a AVA Error
func APIIsEmpty(e error, details interface{}) *errorAVA.Error {
	return APIIsEmptySkip(e, details, errorAVA.RetrieveCallDefault)
}

// APIIsEmptySkip is a AVA Error
func APIIsEmptySkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusAPIIsEmpty,
		Message:  statusTextFunc(statusAPIIsEmpty),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
