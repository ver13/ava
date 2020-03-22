package error

import (
	"fmt"
)

// GroupTypeUnknown is a AVA Error
func GroupTypeUnknown(e error, details interface{}) *Error {
	return GroupTypeUnknownSkip(e, details, RetrieveCallDefault)
}

// GroupTypeUnknownSkip is a AVA Error
func GroupTypeUnknownSkip(e error, details interface{}, skip int) *Error {
	err := Error{
		Group:    GroupUnknown,
		Subgroup: SubgroupUnknown,
		Code:     statusGroupUnknown,
		Message:  statusTextFunc(statusGroupUnknown),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
