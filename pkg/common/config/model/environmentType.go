//go:generate ava-enum -f=$GOFILE --marshal --lower

package model

// EnvironmentType x ENUM(
// Development,
// Test,
// Integration,
// Production,
// Unknown,
// )
type EnvironmentType int32
