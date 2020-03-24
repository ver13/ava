//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package hash

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type HashI interface {
	Version() int
	Hash64(v interface{}) (uint64, *errorAVA.Error)
	HashMD5(c interface{}) string
	HashSHA256(c interface{}) []byte
	Md5(c interface{}) []byte
	Sha1(c interface{}) []byte
	Dump(c interface{}) []byte
}
