package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func URLParseWrong(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    statusURLParseWrongCode,
		Message: statusTextFunc(statusURLParseWrongCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
