//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package url

import (
	"net/url"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// URL defines the interface for all the URL manipulation required by AVA
type URLI interface {
	CleanHosts([]string) []string
	CleanHost(string) (string, *errorAVA.Error)
	CleanPath(string) string
	GetEndpointPath(string, []string) string
	Parse(string) (*url.URL, *errorAVA.Error)
}
