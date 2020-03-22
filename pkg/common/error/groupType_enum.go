// Code generated by go-enum
// DO NOT EDIT!

package error

import (
	"fmt"
	"strings"
)

const (
	// GroupGeneral is a Group of type General
	GroupGeneral Group = iota
	// GroupModel is a Group of type Model
	GroupModel
	// GroupSerializer is a Group of type Serializer
	GroupSerializer
	// GroupEncoder is a Group of type Encoder
	GroupEncoder
	// GroupServer is a Group of type Server
	GroupServer
	// GroupConfig is a Group of type Config
	GroupConfig
	// GroupLogger is a Group of type Logger
	GroupLogger
	// GroupFile is a Group of type File
	GroupFile
	// GroupBlockchain is a Group of type Blockchain
	GroupBlockchain
	// GroupDatabase is a Group of type Database
	GroupDatabase
	// GroupHttp is a Group of type Http
	GroupHttp
	// GroupMicroservice is a Group of type Microservice
	GroupMicroservice
	// GroupMessageCoder is a Group of type MessageCoder
	GroupMessageCoder
	// GroupTime is a Group of type Time
	GroupTime
	// GroupApiTime is a Group of type ApiTime
	GroupApiTime
	// GroupTransport is a Group of type Transport
	GroupTransport
	// GroupCompress is a Group of type Compress
	GroupCompress
	// GroupIO is a Group of type IO
	GroupIO
	// GroupCrypto is a Group of type Crypto
	GroupCrypto
	// GroupQR is a Group of type QR
	GroupQR
	// GroupValidator is a Group of type Validator
	GroupValidator
	// GroupString is a Group of type String
	GroupString
	// GroupUtils is a Group of type Utils
	GroupUtils
	// GroupClient is a Group of type Client
	GroupClient
	// GroupGeneratorEnum is a Group of type GeneratorEnum
	GroupGeneratorEnum
	// GroupRouter is a Group of type Router
	GroupRouter
	// GroupUnknown is a Group of type Unknown
	GroupUnknown
)

const _GroupName = "GeneralModelSerializerEncoderServerConfigLoggerFileBlockchainDatabaseHttpMicroserviceMessageCoderTimeApiTimeTransportCompressIOCryptoQRValidatorStringUtilsClientGeneratorEnumRouterUnknown"

var _GroupMap = map[Group]string{
	0:  _GroupName[0:7],
	1:  _GroupName[7:12],
	2:  _GroupName[12:22],
	3:  _GroupName[22:29],
	4:  _GroupName[29:35],
	5:  _GroupName[35:41],
	6:  _GroupName[41:47],
	7:  _GroupName[47:51],
	8:  _GroupName[51:61],
	9:  _GroupName[61:69],
	10: _GroupName[69:73],
	11: _GroupName[73:85],
	12: _GroupName[85:97],
	13: _GroupName[97:101],
	14: _GroupName[101:108],
	15: _GroupName[108:117],
	16: _GroupName[117:125],
	17: _GroupName[125:127],
	18: _GroupName[127:133],
	19: _GroupName[133:135],
	20: _GroupName[135:144],
	21: _GroupName[144:150],
	22: _GroupName[150:155],
	23: _GroupName[155:161],
	24: _GroupName[161:174],
	25: _GroupName[174:180],
	26: _GroupName[180:187],
}

// String implements the Stringer interface.
func (x Group) String() string {
	if str, ok := _GroupMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Group(%d)", x)
}

var _GroupValue = map[string]Group{
	_GroupName[0:7]:                      0,
	strings.ToLower(_GroupName[0:7]):     0,
	_GroupName[7:12]:                     1,
	strings.ToLower(_GroupName[7:12]):    1,
	_GroupName[12:22]:                    2,
	strings.ToLower(_GroupName[12:22]):   2,
	_GroupName[22:29]:                    3,
	strings.ToLower(_GroupName[22:29]):   3,
	_GroupName[29:35]:                    4,
	strings.ToLower(_GroupName[29:35]):   4,
	_GroupName[35:41]:                    5,
	strings.ToLower(_GroupName[35:41]):   5,
	_GroupName[41:47]:                    6,
	strings.ToLower(_GroupName[41:47]):   6,
	_GroupName[47:51]:                    7,
	strings.ToLower(_GroupName[47:51]):   7,
	_GroupName[51:61]:                    8,
	strings.ToLower(_GroupName[51:61]):   8,
	_GroupName[61:69]:                    9,
	strings.ToLower(_GroupName[61:69]):   9,
	_GroupName[69:73]:                    10,
	strings.ToLower(_GroupName[69:73]):   10,
	_GroupName[73:85]:                    11,
	strings.ToLower(_GroupName[73:85]):   11,
	_GroupName[85:97]:                    12,
	strings.ToLower(_GroupName[85:97]):   12,
	_GroupName[97:101]:                   13,
	strings.ToLower(_GroupName[97:101]):  13,
	_GroupName[101:108]:                  14,
	strings.ToLower(_GroupName[101:108]): 14,
	_GroupName[108:117]:                  15,
	strings.ToLower(_GroupName[108:117]): 15,
	_GroupName[117:125]:                  16,
	strings.ToLower(_GroupName[117:125]): 16,
	_GroupName[125:127]:                  17,
	strings.ToLower(_GroupName[125:127]): 17,
	_GroupName[127:133]:                  18,
	strings.ToLower(_GroupName[127:133]): 18,
	_GroupName[133:135]:                  19,
	strings.ToLower(_GroupName[133:135]): 19,
	_GroupName[135:144]:                  20,
	strings.ToLower(_GroupName[135:144]): 20,
	_GroupName[144:150]:                  21,
	strings.ToLower(_GroupName[144:150]): 21,
	_GroupName[150:155]:                  22,
	strings.ToLower(_GroupName[150:155]): 22,
	_GroupName[155:161]:                  23,
	strings.ToLower(_GroupName[155:161]): 23,
	_GroupName[161:174]:                  24,
	strings.ToLower(_GroupName[161:174]): 24,
	_GroupName[174:180]:                  25,
	strings.ToLower(_GroupName[174:180]): 25,
	_GroupName[180:187]:                  26,
	strings.ToLower(_GroupName[180:187]): 26,
}

// ParseGroup attempts to convert a string to a Group
func ParseGroup(name string) (Group, *Error) {
	if x, ok := _GroupValue[name]; ok {
		return x, nil
	}
	return GroupUnknown, GroupTypeUnknownSkip(nil, fmt.Sprintf("%s is not a valid Group", name), 4)
}

// MarshalText implements the text marshaller method
func (x Group) MarshalText() ([]byte, *Error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x Group) UnmarshalText(text []byte) (Group, *Error) {
	name := string(text)
	tmp, err := ParseGroup(name)
	if err != nil {
		return GroupUnknown, err
	}
	x = tmp
	return tmp, nil
}
