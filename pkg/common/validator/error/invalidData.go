package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// InvalidData is a AVA Error
func InvalidData(e error, details interface{}) *errorAVA.Error {
	return InvalidDataSkip(e, details, 3)
}

// InvalidDataSkip is a AVA Error
func InvalidDataSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupValidator,
		Code:     statusInvalidData,
		Message:  statusTextFunc(statusInvalidData),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
