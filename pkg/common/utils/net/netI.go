//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package net

import (
	"net"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type NETI interface {
	HostPort(addr string, port interface{}) string
	Listen(addr string, fn func(string) (net.Listener, *errorAVA.Error)) (net.Listener, *errorAVA.Error)
}
