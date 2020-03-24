package error

import (
	"fmt"
	
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// URLParseWrong is a AVA Error
func URLParseWrong(e error, details interface{}) *errorAVA.Error {
	return URLParseWrongSkip(e, details, 3)
}

// URLParseWrongSkip is a AVA Error
func URLParseWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUtils,
		Subgroup: errorAVA.SubgroupURL,
		Code:     statusURLParseWrong,
		Message:  statusTextFunc(statusURLParseWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
