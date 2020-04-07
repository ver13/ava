package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// SemanticVersionError is a AVA Error
func SemanticVersionError(e error, details interface{}) *errorAVA.Error {
	return SemanticVersionErrorSkip(e, details, errorAVA.RetrieveCallDefault)
}

// SemanticVersionErrorSkip is a AVA Error
func SemanticVersionErrorSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupVersion,
		Code:     statusSemanticVersionError,
		Message:  statusTextFunc(statusSemanticVersionError),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
