package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ParseInLocation is a AVA Error
func ParseInLocation(e error, details interface{}) *errorAVA.Error {
	return ParseInLocationSkip(e, details, 3)
}

// ParseInLocationSkip is a AVA Error
func ParseInLocationSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusParseInLocation,
		Message:  statusTextFunc(statusParseInLocation),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
