package client

import (
	"context"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Client is the interface used to make requests to services.
// It supports Request/Response via Transport and Publishing via the Broker.
// It also supports bidirectional streaming of requests.
type ClientI interface {
	Init(...Option) *errorAVA.Error
	Options() Options
	NewMessage(topic string, msg interface{}, opts ...MessageOption) MessageI
	NewRequest(service, endpoint string, req interface{}, reqOpts ...RequestOption) RequestI
	Call(ctx context.Context, req RequestI, rsp interface{}, opts ...CallOption) *errorAVA.Error
	Stream(ctx context.Context, req RequestI, opts ...CallOption) (StreamI, *errorAVA.Error)
	Publish(ctx context.Context, msg MessageI, opts ...PublishOption) *errorAVA.Error
	String() string
}
