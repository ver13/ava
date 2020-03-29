package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// CurvePreferencesWrong is a AVA Error
func CurvePreferencesWrong(e error, details interface{}) *errorAVA.Error {
	return CurvePreferencesWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// CurvePreferencesWrongSkip is a AVA Error
func CurvePreferencesWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupConfig,
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Code:     statusCurvePreferencesWrong,
		Message:  statusTextFunc(statusCurvePreferencesWrong),
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
