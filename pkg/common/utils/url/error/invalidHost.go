package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// InvalidHost is a AVA Error
func InvalidHost(e error, details interface{}) *errorAVA.Error {
	return InvalidHostSkip(e, details, 3)
}

// InvalidHostSkip is a AVA Error
func InvalidHostSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUtils,
		Subgroup: errorAVA.SubgroupURL,
		Code:     statusInvalidHost,
		Message:  statusTextFunc(statusInvalidHost),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
