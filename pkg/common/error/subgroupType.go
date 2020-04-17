//go:generate go-enum -f=$GOFILE --marshal --lower

package error

// Subgroup x ENUM(
// General
// DiscoveryService
// BrokerService
// CircuitBreakerService
// MetricsService
// Client
// Server
// Selected
// Serializer
// Hash
// QR
// Version
// Config
// Time
// Validator
// String
// Logger
// File
// GRPC
// NET
// URL
// Codec
// Event
// AVAEnum
// Unknown
// )
type Subgroup int32
