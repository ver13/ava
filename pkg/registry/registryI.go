package registry

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type RegistryI interface {
	Init(...Option) *errorAVA.Error
	Options() Options
	Register(*Service, ...RegisterOption) *errorAVA.Error
	Deregister(*Service) *errorAVA.Error
	GetService(string) ([]*Service, *errorAVA.Error)
	ListServices() ([]*Service, *errorAVA.Error)
	Watch(...WatchOption) (WatcherI, *errorAVA.Error)
	String() string
}
