# AVA [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/ver13/ava?tab=doc) [![Travis CI](https://api.travis-ci.org/ver13/ava.svg?branch=master)](https://travis-ci.org/ver13/ava) [![Go Report Card](https://goreportcard.com/badge/ver13/ava)](https://goreportcard.com/report/github.com/ver13/ava)

AVA is a framework for microservice development.

## Overview

AVA provides the core requirements for distributed systems development including RPC and Event driven communication.
The **micro** philosophy is sane defaults with a pluggable architecture. We provide defaults to get you started quickly
but everything can be easily swapped out.

<img src="https://ver13.github.io/ava/docs/images/ava.svg" />

## Features

AVA abstracts away the details of distributed systems. Here are the main features.

- **Core** - AVA Framework Core functionality provides the basis for developing and improving third-party integration. It offers the basic functionality that is needed to connect to a database system, blockchain, RESTful APIConfig, etc.

- **Tools** -

- **Service Discovery** - Automatic service registration and name resolution. Service discovery is at the core of micro service
development. When service A needs to speak to service B it needs the location of that service. The default discovery mechanism is
multicast DNS (mdns), a zeroconf system.

- **Load Balancing** - Client side load balancing built on service discovery. Once we have the addresses of any number of instances
of a service we now need a way to decide which node to route to. We use random hashed load balancing to provide even distribution
across the services and retry a different node if there's a problem.

- **Message Encoding** - Dynamic message encoding based on content-type. The client and server will use codecs along with content-type
to seamlessly encode and decode Go types for you. Any variety of messages could be encoded and sent from different clients. The client
and server handle this by default. This includes protobuf and json by default.

- **Request/Response** - RPC based request/response with support for bidirectional streaming. We provide an abstraction for synchronous
communication. A request made to a service will be automatically resolved, load balanced, dialled and streamed. The default
transport is [gRPC](https://grpc.io/).

- **Async Messaging** - PubSub is built in as a first class citizen for asynchronous communication and event driven architectures.
Event notifications are a core pattern in micro service development. The default messaging system is an embedded [NATS](https://nats.io/)
server.

## Contents

Contents of this repository:

| Directory | Description                                                     |
| --------- | ----------------------------------------------------------------|
| Core      | Logger, Configuration, Version, Time, Serializer, .....         |
| Tools     | avaEnum, avaWrapper                                             |
| Broker    | PubSub messaging; NATS, NSQ, RabbitMQ, Kafka                    |
| Client    | RPC Clients; gRPC, HTTP                                         |
| Codec     | Message Encoding; BSON, Mercury                                 |
| Registry  | Service Discovery; Etcd, Gossip, NATS                           |
| Selector  | Load balancing; Label, Cache, Static                            |
| Server    | RPC Servers; gRPC, HTTP                                         |
| Transport | Bidirectional Streaming; NATS, RabbitMQ                         |
| Wrapper   | Middleware; Circuit Breakers, Rate Limiting, Tracing, Monitoring|

## Getting Started

See the [docs](https://ver13.github.io/ava/docs/framework.html) for detailed information on the architecture, installation and use of AVA.
