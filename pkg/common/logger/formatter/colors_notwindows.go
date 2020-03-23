// +build !windows

package formatter

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorLoggerAVA "github.com/ver13/ava/pkg/common/logger/error"
)

func windowsNativeANSI(_ bool, _ bool, _ interface{}) (bool, *errorAVA.Error) {
	return false, errorLoggerAVA.NotAvailable(nil, "Not available on this platform: Windows.")
}
