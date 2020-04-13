//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package http

import (
	httpModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type CORSConfigI interface {
	ReadLocal(fileName string) (*httpModelConfigAVA.CORS, *errorAVA.Error)
	Parser() (*httpModelConfigAVA.CORS, *errorAVA.Error)
	Serializer(serializerAVA.SerializerType) ([]byte, *errorAVA.Error)
}
