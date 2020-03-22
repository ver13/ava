//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package file

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type FileI interface {
	FileGetBytes(filenameOrURL string, timeout ...time.Duration) ([]byte, *errorAVA.Error)
	FileSetBytes(filename string, data []byte) *errorAVA.Error
	FileAppendBytes(filename string, data []byte) *errorAVA.Error
	FileGetString(filenameOrURL string, timeout ...time.Duration) (string, *errorAVA.Error)
	FileSetString(filename string, data string) *errorAVA.Error
	FileAppendString(filename string, data string) *errorAVA.Error
	FileGetJSON(filenameOrURL string, timeout ...time.Duration) (result interface{}, err *errorAVA.Error)
	FileUnmarshallJSON(filenameOrURL string, result interface{}, timeout ...time.Duration) *errorAVA.Error
	FileSetJSON(filename string, data interface{}) *errorAVA.Error
	FileSetJSONIndent(filename string, data interface{}, indent string) *errorAVA.Error
	FileGetXML(filenameOrURL string, timeout ...time.Duration) (result interface{}, err *errorAVA.Error)
	FileUnmarshallXML(filenameOrURL string, result interface{}, timeout ...time.Duration) *errorAVA.Error
	FileSetXML(filename string, data interface{}) *errorAVA.Error
	FileGetCSV(filenameOrURL string, timeout ...time.Duration) ([][]string, *errorAVA.Error)
	FileSetCSV(filename string, records [][]string) *errorAVA.Error
	FileGetLines(filenameOrURL string, timeout ...time.Duration) (lines []string, err *errorAVA.Error)
	FileSetLines(filename string, lines []string) *errorAVA.Error
	FileGetNonEmptyLines(filenameOrURL string, timeout ...time.Duration) (lines []string, err *errorAVA.Error)
	FileGetLastLine(filenameOrURL string, timeout ...time.Duration) (line string, err *errorAVA.Error)
	FileTimeModified(filename string) time.Time
	FileExists(filename string) *errorAVA.Error
	FileIsDir(dirname string) bool
	FileFind(searchDirs []string, filenames ...string) (filePath string, found bool)
	FileFindModified(searchDirs []string, filenames ...string) (filePath string, found bool, modified time.Time)
	FileTouch(filename string) *errorAVA.Error
	FileMD5String(filenameOrURL string) (string, *errorAVA.Error)
	FileMD5Bytes(filenameOrURL string) ([]byte, *errorAVA.Error)
	FileCRC64(filenameOrURL string) (uint64, *errorAVA.Error)
	FileGetInflate(filenameOrURL string) ([]byte, *errorAVA.Error)
	FileGetGz(filenameOrURL string) ([]byte, *errorAVA.Error)
	FileSize(filename string) int64
	FilePrintf(filename, format string, args ...interface{}) *errorAVA.Error
	FileAppendPrintf(filename, format string, args ...interface{}) *errorAVA.Error
	FileScanf(filename, format string, args ...interface{}) *errorAVA.Error
	ListDir(dir string) ([]string, *errorAVA.Error)
	ListDirFiles(dir string) ([]string, *errorAVA.Error)
	ListDirDirectories(dir string) ([]string, *errorAVA.Error)
	FileCopy(source string, dest string) (err *errorAVA.Error)
	FileCopyDir(source string, dest string) (err *errorAVA.Error)
	Abs(path string) (string, *errorAVA.Error)
}
