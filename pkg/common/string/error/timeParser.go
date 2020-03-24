package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// TimeParser is a AVA Error
func TimeParser(e error, details interface{}) *errorAVA.Error {
	return TimeParserSkip(e, details, errorAVA.RetrieveCallDefault)
}

// TimeParserSkip is a AVA Error
func TimeParserSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupString,
		Code:     statusTimeParser,
		Message:  statusTextFunc(statusTimeParser),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
