package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// DiscoveryServiceUnknown is a AVA Error
func DiscoveryServiceUnknown(e error, details interface{}) *errorAVA.Error {
	return DiscoveryServiceUnknownSkip(e, details, errorAVA.RetrieveCallDefault)
}

// DiscoveryServiceUnknownSkip is a AVA Error
func DiscoveryServiceUnknownSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusDiscoveryServiceUnknown,
		Message:  statusTextFunc(statusDiscoveryServiceUnknown),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
