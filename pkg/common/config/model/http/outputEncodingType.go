//go:generate go-enum -f=$GOFILE --marshal --lower

package http

// OutputEncodingType x ENUM(
// JSON
// XML
// STRING
// NOOP
// UNKNOWN
// )
type OutputEncodingType int32
