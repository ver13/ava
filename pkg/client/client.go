package client

import (
	"context"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

var (
	// DefaultClient is a default client to use out of the box
	DefaultClient ClientI = newRpcClient()
	// DefaultBackoff is the default backoff function for retries
	DefaultBackoff = exponentialBackoff
	// DefaultRetry is the default check-for-retry function for retries
	DefaultRetry = RetryOnError
	// DefaultRetries is the default number of times a request is tried
	DefaultRetries = 1
	// DefaultRequestTimeout is the default request timeout
	DefaultRequestTimeout = time.Second * 5
	// DefaultPoolSize sets the connection pool size
	DefaultPoolSize = 100
	// DefaultPoolTTL sets the connection pool ttl
	DefaultPoolTTL = time.Minute

	// NewClient returns a new client
	NewClient func(...Option) ClientI = newRpcClient
)

// Makes a synchronous call to a service using the default client
func Call(ctx context.Context, request RequestI, response interface{}, opts ...CallOption) *errorAVA.Error {
	return DefaultClient.Call(ctx, request, response, opts...)
}

// Publishes a publication using the default client. Using the underlying broker
// set within the options.
func Publish(ctx context.Context, msg MessageI, opts ...PublishOption) *errorAVA.Error {
	return DefaultClient.Publish(ctx, msg, opts...)
}

// Creates a new message using the default client
func NewMessage(topic string, payload interface{}, opts ...MessageOption) MessageI {
	return DefaultClient.NewMessage(topic, payload, opts...)
}

// Creates a new request using the default client. Content Type will
// be set to the default within options and use the appropriate codec
func NewRequest(service, endpoint string, request interface{}, reqOpts ...RequestOption) RequestI {
	return DefaultClient.NewRequest(service, endpoint, request, reqOpts...)
}

// Creates a streaming connection with a service and returns responses on the
// channel passed in. It's up to the user to close the streamer.
func NewStream(ctx context.Context, request RequestI, opts ...CallOption) (StreamI, *errorAVA.Error) {
	return DefaultClient.Stream(ctx, request, opts...)
}

func String() string {
	return DefaultClient.String()
}
