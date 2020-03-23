//go:generate avaEnum -f=$GOFILE --marshal --lower

package logger

// LogFormatterType x ENUM(
// Text
// JSON
// ApacheCommonLog
// ApacheCombinedLog
// ApacheErrorLog
// RFC3164Log
// RFC5424Log
// CommonLogFileFormat
// Unknown
// )
type LogFormatterType int32
