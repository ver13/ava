package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FailedGetInterfaces is a AVA Error
func FailedGetInterfaces(e error, details interface{}) *errorAVA.Error {
	return FailedGetInterfacesSkip(e, details, 3)
}

// FailedGetInterfacesSkip is a AVA Error
func FailedGetInterfacesSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusFailedGetInterfaces,
		Message:  statusTextFunc(statusFailedGetInterfaces),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
