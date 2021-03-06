// Code generated by go-enum
// DO NOT EDIT!

package brokerService

import (
	"fmt"
	"strings"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

const (
	// BrokerServiceTypeHTTP is a BrokerServiceType of type HTTP
	BrokerServiceTypeHTTP BrokerServiceType = iota
	// BrokerServiceTypeKafka is a BrokerServiceType of type Kafka
	BrokerServiceTypeKafka
	// BrokerServiceTypeMemory is a BrokerServiceType of type Memory
	BrokerServiceTypeMemory
	// BrokerServiceTypeNATS is a BrokerServiceType of type NATS
	BrokerServiceTypeNATS
	// BrokerServiceTypeRabbitMQ is a BrokerServiceType of type RabbitMQ
	BrokerServiceTypeRabbitMQ
	// BrokerServiceTypeUnknown is a BrokerServiceType of type Unknown
	BrokerServiceTypeUnknown
)

const _BrokerServiceTypeName = "HTTPKafkaMemoryNATSRabbitMQUnknown"

var _BrokerServiceTypeMap = map[BrokerServiceType]string{
	0: _BrokerServiceTypeName[0:4],
	1: _BrokerServiceTypeName[4:9],
	2: _BrokerServiceTypeName[9:15],
	3: _BrokerServiceTypeName[15:19],
	4: _BrokerServiceTypeName[19:27],
	5: _BrokerServiceTypeName[27:34],
}

// String implements the Stringer interface.
func (x BrokerServiceType) String() string {
	if str, ok := _BrokerServiceTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("BrokerServiceType(%d)", x)
}

var _BrokerServiceTypeValue = map[string]BrokerServiceType{
	_BrokerServiceTypeName[0:4]:                    0,
	strings.ToLower(_BrokerServiceTypeName[0:4]):   0,
	_BrokerServiceTypeName[4:9]:                    1,
	strings.ToLower(_BrokerServiceTypeName[4:9]):   1,
	_BrokerServiceTypeName[9:15]:                   2,
	strings.ToLower(_BrokerServiceTypeName[9:15]):  2,
	_BrokerServiceTypeName[15:19]:                  3,
	strings.ToLower(_BrokerServiceTypeName[15:19]): 3,
	_BrokerServiceTypeName[19:27]:                  4,
	strings.ToLower(_BrokerServiceTypeName[19:27]): 4,
	_BrokerServiceTypeName[27:34]:                  5,
	strings.ToLower(_BrokerServiceTypeName[27:34]): 5,
}

// ParseBrokerServiceType attempts to convert a string to a BrokerServiceType
func ParseBrokerServiceType(name string) (BrokerServiceType, *errorAVA.Error) {
	if x, ok := _BrokerServiceTypeValue[name]; ok {
		return x, nil
	}
	return BrokerServiceTypeUnknown, errorConfigAVA.BrokerServiceUnknown(nil, fmt.Sprintf("%s is not a valid broker service type", name))
}

// MarshalText implements the text marshaller method
func (x BrokerServiceType) MarshalText() ([]byte, *errorAVA.Error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x BrokerServiceType) UnmarshalText(text []byte) (BrokerServiceType, *errorAVA.Error) {
	name := string(text)
	tmp, err := ParseBrokerServiceType(name)
	if err != nil {
		return BrokerServiceTypeUnknown, err
	}
	x = tmp
	return tmp, nil
}
