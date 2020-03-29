package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// BrokerServiceUnknown is a AVA Error
func BrokerServiceUnknown(e error, details interface{}) *errorAVA.Error {
	return BrokerServiceUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// BrokerServiceUnknownSkip is a AVA Error
func BrokerServiceUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusBrokerServiceUnknown,
		Message:  statusTextFunc(statusBrokerServiceUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
