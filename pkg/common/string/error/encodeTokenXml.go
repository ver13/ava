package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// EncodeTokenXml is a AVA Error
func EncodeTokenXml(e error, details interface{}) *errorAVA.Error {
	return EncodeTokenXmlSkip(e, details, errorAVA.RetrieveCallDefault)
}

// EncodeTokenXmlSkip is a AVA Error
func EncodeTokenXmlSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupString,
		Code:     statusEncodeTokenXml,
		Message:  statusTextFunc(statusEncodeTokenXml),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
