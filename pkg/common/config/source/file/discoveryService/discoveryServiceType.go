//go:generate go-enum -f=$GOFILE --marshal --lower

package discoveryService

// DiscoveryServiceType x ENUM(
// Eureka
// Consul
// Memory
// Kubernetes
// Etcd
// Gossip
// NATS
// Zookeeper
// Unknown
// )
type DiscoveryServiceType int32
