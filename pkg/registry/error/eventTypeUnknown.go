package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// EventTypeUnknown is a AVA Error
func EventTypeUnknown(e error, details interface{}) *errorAVA.Error {
	return EventTypeUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// EventTypeUnknownSkip is a AVA Error
func EventTypeUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupRegistry,
		Subgroup: errorAVA.SubgroupEvent,
		Code:     statusEventTypeUnknown,
		Message:  statusTextFunc(statusEventTypeUnknown),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
