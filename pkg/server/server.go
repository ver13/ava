package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	uuidUtilsAVA "github.com/ver13/ava/pkg/common/utils/uuid"
)

var (
	DefaultAddress                  = ":0"
	DefaultName                     = "go.micro.server"
	DefaultVersion                  = "latest"
	DefaultId                       = uuidUtilsAVA.GetInstance().NewUUID()
	DefaultServer           ServerI = newRpcServer()
	DefaultRouter                   = newRpcRouter()
	DefaultRegisterCheck            = func(context.Context) *errorAVA.Error { return nil }
	DefaultRegisterInterval         = time.Second * 30
	DefaultRegisterTTL              = time.Minute

	// NewServer creates a new server
	NewServer func(...Option) ServerI = newRpcServer
)

// Init initialises the default server with options passed in
func Init(opt ...Option) {
	if DefaultServer == nil {
		DefaultServer = newRpcServer(opt...)
	}
	DefaultServer.Init(opt...)
}

// NewRouter returns a new router
func NewRouter() *router {
	return newRpcRouter()
}

// NewSubscriber creates a new subscriber interface with the given topic
// and handler using the default server
func NewSubscriber(topic string, h interface{}, opts ...SubscriberOption) SubscriberI {
	return DefaultServer.NewSubscriber(topic, h, opts...)
}

// NewHandler creates a new handler interface using the default server
// Handlers are required to be a public object with public
// endpoints. Call to a service endpoint such as Foo.Bar expects
// the type:
//
//	type Foo struct {}
//	func (f *Foo) Bar(ctx, req, rsp) error {
//		return nil
//	}
//
func NewHandler(h interface{}, opts ...HandlerOption) HandlerI {
	return DefaultServer.NewHandler(h, opts...)
}

// Handle registers a handler interface with the default server to
// handle inbound requests
func Handle(h HandlerI) *errorAVA.Error {
	return DefaultServer.Handle(h)
}

// Subscribe registers a subscriber interface with the default server
// which subscribes to specified topic with the broker
func Subscribe(s SubscriberI) *errorAVA.Error {
	return DefaultServer.Subscribe(s)
}

// Run starts the default server and waits for a kill
// signal before exiting. Also registers/deregisters the server
func Run() *errorAVA.Error {
	if err := Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	loggerAVA.Infof("Received signal %s", <-ch)

	return Stop()
}

// Start starts the default server
func Start() *errorAVA.Error {
	config := DefaultServer.Options()
	loggerAVA.Infof("Starting server %s id %s", config.Name, config.Id)

	return DefaultServer.Start()
}

// Stop stops the default server
func Stop() *errorAVA.Error {
	loggerAVA.Infof("Stopping server")

	return DefaultServer.Stop()
}

// String returns name of Server implementation
func String() string {
	return DefaultServer.String()
}
