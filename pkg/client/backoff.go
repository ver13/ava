package client

import (
	"context"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type BackoffFunc func(ctx context.Context, req RequestI, attempts int) (time.Duration, *errorAVA.Error)

func exponentialBackoff(ctx context.Context, req RequestI, attempts int) (time.Duration, *errorAVA.Error) {
	return backoff.Do(attempts), nil
}
