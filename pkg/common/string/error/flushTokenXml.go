package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// FlushTokenXml is a AVA Error
func FlushTokenXml(e error, details interface{}) *errorAVA.Error {
	return FlushTokenXmlSkip(e, details, errorAVA.RetrieveCallDefault)
}

// FlushTokenXmlSkip is a AVA Error
func FlushTokenXmlSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupString,
		Code:     statusFlushTokenXml,
		Message:  statusTextFunc(statusFlushTokenXml),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
