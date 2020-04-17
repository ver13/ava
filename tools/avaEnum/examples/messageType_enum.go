// Code generated by go-enum
// DO NOT EDIT!

package examples

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

const (
	// Error is a MessageType of type Error
	Error MessageType = iota
	// Request is a MessageType of type Request
	Request
	// Response is a MessageType of type Response
	Response
	// Event is a MessageType of type Event
	Event
	// Unknown is a MessageType of type Unknown
	Unknown
)

const _MessageTypeName = "ErrorRequestResponseEventUnknown"

var _MessageTypeNames = []string{
	_MessageTypeName[0:5],
	_MessageTypeName[5:12],
	_MessageTypeName[12:20],
	_MessageTypeName[20:25],
	_MessageTypeName[25:32],
}

// MessageTypeNames returns a list of possible string values of MessageType.
func MessageTypeNames() []string {
	tmp := make([]string, len(_MessageTypeNames))
	copy(tmp, _MessageTypeNames)
	return tmp
}

var _MessageTypeMap = map[MessageType]string{
	0: _MessageTypeName[0:5],
	1: _MessageTypeName[5:12],
	2: _MessageTypeName[12:20],
	3: _MessageTypeName[20:25],
	4: _MessageTypeName[25:32],
}

// String implements the Stringer interface.
func (x MessageType) String() string {
	if str, ok := _MessageTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("MessageType(%d)", x)
}

var _MessageTypeValue = map[string]MessageType{
	_MessageTypeName[0:5]:                    0,
	strings.ToLower(_MessageTypeName[0:5]):   0,
	_MessageTypeName[5:12]:                   1,
	strings.ToLower(_MessageTypeName[5:12]):  1,
	_MessageTypeName[12:20]:                  2,
	strings.ToLower(_MessageTypeName[12:20]): 2,
	_MessageTypeName[20:25]:                  3,
	strings.ToLower(_MessageTypeName[20:25]): 3,
	_MessageTypeName[25:32]:                  4,
	strings.ToLower(_MessageTypeName[25:32]): 4,
}

// ParseMessageType attempts to convert a string to a MessageType
func ParseMessageType(name string) (MessageType, error) {
	if x, ok := _MessageTypeValue[name]; ok {
		return x, nil
	}
	return MessageType(0), fmt.Errorf("%s is not a valid MessageType, try [%s]", name, strings.Join(_MessageTypeNames, ", "))
}

// MarshalText implements the text marshaller method
func (x MessageType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *MessageType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseMessageType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Scan implements the Scanner interface.
func (x *MessageType) Scan(value interface{}) error {
	var name string

	switch v := value.(type) {
	case string:
		name = v
	case []byte:
		name = string(v)
	case nil:
		*x = MessageType(0)
		return nil
	}

	tmp, err := ParseMessageType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Value implements the driver Valuer interface.
func (x MessageType) Value() (driver.Value, error) {
	return x.String(), nil
}

// Set implements the Golang flag.Value interface func.
func (x *MessageType) Set(val string) error {
	v, err := ParseMessageType(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *MessageType) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *MessageType) Type() string {
	return "MessageType"
}
