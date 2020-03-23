package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func FailedGetInterfaces(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    StatusFailedGetInterfacesCode,
		Message: StatusText(StatusFailedGetInterfacesCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
