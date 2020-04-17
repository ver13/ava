package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ParseFileWrong is a AVA Error
func ParseFileWrong(e error, details interface{}) *errorAVA.Error {
	return ParseFileWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// ParseFileWrongSkip is a AVA Error
func ParseFileWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupTools,
		Subgroup: errorAVA.SubgroupAVAEnum,
		Code:     statusParseFileWrong,
		Message:  statusTextFunc(statusParseFileWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
