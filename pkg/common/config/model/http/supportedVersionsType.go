//go:generate go-enum -f=$GOFILE --marshal --lower
// +build !test

package http

// SupportedVersionsType x ENUM(
// TLS10 = 769
// TLS11 = 770
// TLS12 = 771
// TLS13 = 772
// )
type SupportedVersionsType uint16
