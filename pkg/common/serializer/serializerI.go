//go:generate gmfGenerateWrap gen -f=${GOFILE} -t implGmf.tmpl -o ${GOFILE}Impl.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t loggerGmf.tmpl -o ${GOFILE}Logger.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t circuitBreakerGmf.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate gmfGenerateTest -f=${GOFILE}

package serializer

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type SerializerI interface {
	Serializer(interface{}) ([]byte, *errorAVA.Error)
	Deserializer([]byte, interface{}) *errorAVA.Error
	String() string
	Type() SerializerType
}
