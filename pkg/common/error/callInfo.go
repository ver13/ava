package error

import (
	"path"
	"runtime"
	"strings"
)

const RetrieveCallDefault = 3

type CallInfo struct {
	PackageName string
	FileName    string
	FuncName    string
	Line        int
}

// retrieves the stack trace of the current call in runtime with 3 skip step
func RetrieveCallInfo() *CallInfo {
	return RetrieveCallInfoSkip(RetrieveCallDefault)
}

// retrieves the stack trace of the current call in runtime with n skip step
func RetrieveCallInfoSkip(skip int) *CallInfo {
	pc, file, line, _ := runtime.Caller(skip)
	_, fileName := path.Split(file)
	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	funcName = parts[2]
	packageName := parts[0] + "." + parts[1]

	return &CallInfo{
		PackageName: packageName,
		FileName:    fileName,
		FuncName:    funcName,
		Line:        line,
	}
}
