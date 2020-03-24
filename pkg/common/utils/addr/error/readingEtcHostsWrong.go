package error

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// ReadingEtcHostsWrong is a AVA Error
func ReadingEtcHostsWrong(e error, details interface{}) *errorAVA.Error {
	return ReadingEtcHostsWrongSkip(e, details, 3)
}

// ReadingEtcHostsWrongSkip is a AVA Error
func ReadingEtcHostsWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupGeneral,
		Subgroup: errorAVA.SubgroupTime,
		Code:     statusReadingEtcHostsWrong,
		Message:  statusTextFunc(statusReadingEtcHostsWrong),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
