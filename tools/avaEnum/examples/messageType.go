//go:generate avaEnum -f=$GOFILE --marshal --flag --sql  --nocamel --lower --names --noprefix --prefix

package examples

// MessageType x ENUM(
// Error
// Request
// Response
// Event
// Unknown
// )
type MessageType int32
