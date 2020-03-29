//go:generate go-enum -f=$GOFILE --marshal --lower
// +build !test

package http

// SupportedCurvesType x ENUM(
// CurveP256 = 23
// CurveP384 = 24
// CurveP521 = 25
// X25519 = 29
// )
type SupportedCurvesType int32
