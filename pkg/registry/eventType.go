//go:generate avaEnum -f=$GOFILE --marshal --lower

package registry

// EventType x ENUM(
// Create     // Create is emitted when a new service is registered
// Delete     // Delete is emitted when an existing service is deregsitered
// Update     // Update is emitted when an existing servicec is updated
// Unknown
// )
type EventType int32
