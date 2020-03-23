//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package string

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type StringI interface {
	MarshalJSON(data interface{}, indent string) string
	ListContains(l []string, s string) bool
	ListContainsCaseInsensitive(l []string, s string) bool
	PrettifyJSON(compactJSON string) string
	EscapeJSON(jsonString string) string
	StripHTMLTags(text string) (plainText string)
	ReplaceHTMLTags(text, replacement string) (plainText string)
	MD5Hex(data string) string
	SHA1Base64(data string) string
	AddURLParam(url, name, value string) string
	ConvertTime(timeString string, formatIn string, formatOut string) (string, *errorAVA.Error)
	CSV(records [][]string) string
	ToInt(s string) int
	ToFloat(s string) float64
	ToBool(s string) bool
	InSlice(s string, slice []string) bool
	JoinFormat(format string, values interface{}, sep string) string
	Join(values interface{}, sep string) string
	FormatBigInt(mem uint64) string
	FormatMemory(mem uint64) string
	ReplaceMulti(str string, fromTo ...string) string
	ToUpperCamelCase(str string) string
	ToLowerCamelCase(str string) string
	ToLower(str string) string
	ToUpper(str string) string
	MapSortedKeys(m map[string]string) []string
	MapGroupedNumberPostfixSortedKeys(m map[string]string) []string
	MapGroupedNumberPostfixSortedValues(m map[string]string) []string
	EndsWithNumber(s string) bool
	SplitNumberPostfix(s string) (base, number string)
	SplitOnce(s, sep string) (pre, post string)
	SplitOnceChar(s string, sep byte) (pre, post string)
	SplitOnceRune(s string, sep rune) (pre, post string)
	MapFunc(f func(string) string, data []string) []string
	Filter(f func(string) bool, data []string) []string
	FindBetween(s, start, stop string) (between, remainder string, found bool)
	Find(s, token string) (remainder string, found bool)
}
