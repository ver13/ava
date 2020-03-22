//go:generate ava-enum -f=$GOFILE --marshal --lower

package error

// Group x ENUM(
// General
// Model
// Serializer
// Encoder
// Server
// Config
// Logger
// File
// Blockchain
// Database
// Http
// Microservice
// MessageCoder
// Time
// ApiTime
// Transport
// Compress
// IO
// Crypto
// QR
// Validator
// String
// Utils
// Client
// GeneratorEnum
// Router
// Unknown
// )
type Group int32