//go:generate gmfGenerateWrap gen -f=${GOFILE} -t implGmf.tmpl -o ${GOFILE}Impl.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t loggerGmf.tmpl -o ${GOFILE}Logger.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t circuitBreakerGmf.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate gmfGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate gmfGenerateTest -f=${GOFILE}

package uri

import (
	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

// URIParser defines the interface for all the URI manipulation required by Gmf
type URIParserI interface {
	CleanHosts([]string) []string
	CleanHost(string) (string, errorGmf.ErrorGmfI)
	CleanPath(string) string
	GetEndpointPath(string, []string) string
	Parse(string) (*url, errorGmf.ErrorGmfI)
}
