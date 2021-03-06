// Code generated by go-enum
// DO NOT EDIT!

package model

import (
	"fmt"
	"strings"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

const (
	// EnvironmentTypeDevelopment is a EnvironmentType of type Development
	EnvironmentTypeDevelopment EnvironmentType = iota
	// EnvironmentTypeTest is a EnvironmentType of type Test
	EnvironmentTypeTest
	// EnvironmentTypeIntegration is a EnvironmentType of type Integration
	EnvironmentTypeIntegration
	// EnvironmentTypeProduction is a EnvironmentType of type Production
	EnvironmentTypeProduction
	// EnvironmentTypeUnknown is a EnvironmentType of type Unknown
	EnvironmentTypeUnknown
)

const _EnvironmentTypeName = "DevelopmentTestIntegrationProductionUnknown"

var _EnvironmentTypeMap = map[EnvironmentType]string{
	0: _EnvironmentTypeName[0:11],
	1: _EnvironmentTypeName[11:15],
	2: _EnvironmentTypeName[15:26],
	3: _EnvironmentTypeName[26:36],
	4: _EnvironmentTypeName[36:43],
}

// String implements the Stringer interface.
func (x EnvironmentType) String() string {
	if str, ok := _EnvironmentTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("EnvironmentType(%d)", x)
}

var _EnvironmentTypeValue = map[string]EnvironmentType{
	_EnvironmentTypeName[0:11]:                   0,
	strings.ToLower(_EnvironmentTypeName[0:11]):  0,
	_EnvironmentTypeName[11:15]:                  1,
	strings.ToLower(_EnvironmentTypeName[11:15]): 1,
	_EnvironmentTypeName[15:26]:                  2,
	strings.ToLower(_EnvironmentTypeName[15:26]): 2,
	_EnvironmentTypeName[26:36]:                  3,
	strings.ToLower(_EnvironmentTypeName[26:36]): 3,
	_EnvironmentTypeName[36:43]:                  4,
	strings.ToLower(_EnvironmentTypeName[36:43]): 4,
}

// ParseEnvironmentType attempts to convert a string to a EnvironmentType
func ParseEnvironmentType(name string) (EnvironmentType, *errorAVA.Error) {
	if x, ok := _EnvironmentTypeValue[name]; ok {
		return x, nil
	}
	return EnvironmentTypeDevelopment, errorConfigAVA.EnvironmentUnknown(nil, fmt.Errorf("%s is not a valid Environment", name))
}

// MarshalText implements the text marshaller method
func (x EnvironmentType) MarshalText() ([]byte, *errorAVA.Error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x EnvironmentType) UnmarshalText(text []byte) (EnvironmentType, *errorAVA.Error) {
	name := string(text)
	tmp, err := ParseEnvironmentType(name)
	if err != nil {
		return EnvironmentTypeUnknown, err
	}
	x = tmp
	return tmp, nil
}
