package error

import (
	"fmt"
)

// SubgroupTypeUnknown is a AVA Error
func SubgroupTypeUnknown(e error, details interface{}) *Error {
	return SubgroupTypeUnknownSkip(e, details, RetrieveCallDefault)
}

// SubgroupTypeUnknownSkip is a AVA Error
func SubgroupTypeUnknownSkip(e error, details interface{}, skip int) *Error {
	err := Error{
		Group:    GroupUnknown,
		Subgroup: SubgroupUnknown,
		Code:     SubgroupUnknownCode,
		Message:  StatusTextFunc(SubgroupUnknownCode),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
