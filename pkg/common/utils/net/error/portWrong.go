package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func PortWrong(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    StatusPortWrongCode,
		Message: StatusText(StatusPortWrongCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
