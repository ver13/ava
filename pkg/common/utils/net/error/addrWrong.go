package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func AddrWrong(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    StatusAddrWrongCode,
		Message: StatusText(StatusAddrWrongCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
