//go:generate ava-enum -f=$GOFILE --marshal --lower

package brokerService

// BrokerServiceType x ENUM(
// HTTP
// Kafka
// Memory
// NATS
// RabbitMQ
// Unknown
// )
type BrokerServiceType int32
