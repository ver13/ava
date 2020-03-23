//go:generate gmfGenerateWrap gen -f=${GOFILE} -t implGmf.tmpl -o ${GOFILE}Impl.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t loggerGmf.tmpl -o ${GOFILE}Logger.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t circuitBreakerGmf.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate gmfGenerateTest -f=${GOFILE}

package addr

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

type AddrI interface {
	GetIPWithPrefix(prefix string) string
	GetIP() string
	ResolveIPFromHostsFile() (string, errorGmf.ErrorGmfI)
	IPs() []string
	Extract(addr string) (string, errorGmf.ErrorGmfI)
	IsPrivateIP(ipAddr string) bool
}
