//go:generate ava-enum -f=$GOFILE --marshal --lower

package version

// WildcardType x ENUM(
// ENUM(
// None
// Major
// Minor
// Patch
// Unknown
// )
type WildcardType int32
