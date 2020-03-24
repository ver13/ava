//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package addr

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type AddrI interface {
	GetIPWithPrefix(prefix string) string
	GetIP() string
	ResolveIPFromHostsFile() (string, *errorAVA.Error)
	IPs() []string
	Extract(addr string) (string, *errorAVA.Error)
	IsPrivateIP(ipAddr string) bool
}
