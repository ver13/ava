//go:generate gmfGenerateWrap gen -f=${GOFILE} -t implGmf.tmpl -o ${GOFILE}Impl.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t loggerGmf.tmpl -o ${GOFILE}Logger.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t circuitBreakerGmf.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate gmfGenerateTest -f=${GOFILE}

package socket

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
	transportHTTPGmf "github.com/ValentinEncinasRojas/ava/pkg/transport"
)

type SocketI interface {
	Accept(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI
	Process(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI
	Send(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI
	Recv(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI
	Close() errorGmf.ErrorGmfI
}
