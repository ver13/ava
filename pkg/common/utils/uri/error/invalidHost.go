package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func InvalidHost(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    statusInvalidHostCode,
		Message: StatusText(statusInvalidHostCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
