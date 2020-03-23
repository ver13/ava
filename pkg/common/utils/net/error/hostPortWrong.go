package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func HostPortWrong(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    StatusHostPortWrongCode,
		Message: StatusText(StatusHostPortWrongCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
