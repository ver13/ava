// Package proto contains utility functions for working with protobufs
package proto

import (
	routerGmf "github.com/ValentinEncinasRojas/ava/pkg/router"
	protoServiceGmf "github.com/ValentinEncinasRojas/ava/pkg/router/service/proto"
)

// RouteToProto encodes route into protobuf and returns it
func RouteToProto(route routerGmf.Route) *protoServiceGmf.Route {
	return &protoServiceGmf.Route{
		Service: route.Service,
		Address: route.Address,
		Gateway: route.Gateway,
		Network: route.Network,
		Router:  route.Router,
		Link:    route.Link,
		Metric:  int64(route.Metric),
	}
}
