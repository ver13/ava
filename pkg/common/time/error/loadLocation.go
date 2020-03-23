package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// LoadLocation is a AVA Error
func LoadLocation(e error, details interface{}) *errorAVA.Error {
	return LoadLocationSkip(e, details, 3)
}

// LoadLocationSkip is a AVA Error
func LoadLocationSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusLoadLocation,
		Message:  statusTextFunc(statusLoadLocation),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
