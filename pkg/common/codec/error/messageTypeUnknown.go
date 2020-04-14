package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// MessageTypeUnknown is a AVA Error
func MessageTypeUnknown(e error, details interface{}) *errorAVA.Error {
	return MessageTypeUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// MessageTypeUnknownSkip is a AVA Error
func MessageTypeUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupCodec,
		Code:     statusMessageUnknown,
		Message:  statusTextFunc(statusMessageUnknown),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
