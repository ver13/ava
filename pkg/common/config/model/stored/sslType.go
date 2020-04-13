//go:generate avaEnum -f=$GOFILE --marshal --lower

package stored

// SSLType x ENUM(
// Enable
// Disable
// Unknown
// )
type SSLType int32
