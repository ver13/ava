package error

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

func ReadingEtcHostsWrong(e error, details string) errorGmf.ErrorGmfI {
	err := errorGmf.ErrorGmf{
		Group:   errorGmf.GroupTypeUtils,
		Code:    StatusReadingEtcHostsWrongCode,
		Message: StatusText(StatusReadingEtcHostsWrongCode),
		Details: details,
		Err:     e,
		Info:    errorGmf.RetrieveCallInfo(),
	}
	err.Println()
	return &err
}
