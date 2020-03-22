package general

import (
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// DeepCopyWrong is a AVA Error
func DeepCopyWrong(e error, details interface{}) *errorAVA.Error {
	return DeepCopyWrongSkip(e, details, errorAVA.RetrieveCallDefault)
}

// DeepCopyWrongSkip is a AVA Error
func DeepCopyWrongSkip(e error, details interface{}, skip int) *errorAVA.Error {
	err := errorAVA.Error{
		Group:    errorAVA.GroupUnknown,
		Subgroup: errorAVA.SubgroupUnknown,
		Code:     DeepCopyWrongCode,
		Message:  StatusTextFunc(DeepCopyWrongCode),
		Details:  fmt.Sprintf("%v.", details),
		Err:      e,
		Info:     errorAVA.RetrieveCallInfoSkip(skip),
	}
	err.Println()
	return &err
}
