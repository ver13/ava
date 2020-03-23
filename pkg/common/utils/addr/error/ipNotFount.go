package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func IPNotFount(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    StatusIPNotFountCode,
		Message: StatusText(StatusIPNotFountCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
