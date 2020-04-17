package client

import (
	"context"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// note that returning either false or a non-nil error will result in the call not being retried
type RetryFunc func(ctx context.Context, req RequestI, retryCount int, err *errorAVA.Error) (bool, *errorAVA.Error)

// RetryAlways always retry on error
func RetryAlways(ctx context.Context, req RequestI, retryCount int, err error) (bool, *errorAVA.Error) {
	return true, nil
}

// RetryOnError retries a request on a 500 or timeout error
func RetryOnError(ctx context.Context, req RequestI, retryCount int, err *errorAVA.Error) (bool, *errorAVA.Error) {
	if err == nil {
		return false, nil
	}

	e := errors.Parse(err.Error())
	if e == nil {
		return false, nil
	}

	switch e.Code {
	// retry on timeout or internal server error
	case 408, 500:
		return true, nil
	default:
		return false, nil
	}
}
