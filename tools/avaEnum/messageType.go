//go:generate avaEnum -f=$GOFILE --marshal --lower

package codec

// MessageType x ENUM(
// Error
// Request
// Response
// Event
// Unknown
// )
type MessageType int32
