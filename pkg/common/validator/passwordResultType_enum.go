// code generated by go-enum
// DO NOT EDIT!

package validator

import (
	"fmt"
	"strings"
)

const (
	// PasswordResultOK is a PasswordResultType of type PasswordResultOK
	// PasswordResultOK Means the checking ran alright
	PasswordResultOK PasswordResultType = iota
	// PasswordResultDivergent is a PasswordResultType of type PasswordResultDivergent
	// PasswordResultDivergent Password is different from confirmation
	PasswordResultDivergent
	// PasswordResultTooShort is a PasswordResultType of type PasswordResultTooShort
	// PasswordResultTooShort Password is too short
	PasswordResultTooShort
	// PasswordResultTooSimple is a PasswordResultType of type PasswordResultTooSimple
	// PasswordResultTooSimple Given string doesn't satisfy complexity rules
	PasswordResultTooSimple
)

const _PasswordResultTypeName = "PasswordResultOKPasswordResultDivergentPasswordResultTooShortPasswordResultTooSimple"

var _PasswordResultTypeMap = map[PasswordResultType]string{
	0: _PasswordResultTypeName[0:16],
	1: _PasswordResultTypeName[16:39],
	2: _PasswordResultTypeName[39:61],
	3: _PasswordResultTypeName[61:84],
}

// String implements the Stringer interface.
func (x PasswordResultType) String() string {
	if str, ok := _PasswordResultTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PasswordResultType(%d)", x)
}

var _PasswordResultTypeValue = map[string]PasswordResultType{
	_PasswordResultTypeName[0:16]:                   0,
	strings.ToLower(_PasswordResultTypeName[0:16]):  0,
	_PasswordResultTypeName[16:39]:                  1,
	strings.ToLower(_PasswordResultTypeName[16:39]): 1,
	_PasswordResultTypeName[39:61]:                  2,
	strings.ToLower(_PasswordResultTypeName[39:61]): 2,
	_PasswordResultTypeName[61:84]:                  3,
	strings.ToLower(_PasswordResultTypeName[61:84]): 3,
}

// ParsePasswordResultType attempts to convert a string to a PasswordResultType
func ParsePasswordResultType(name string) (PasswordResultType, error) {
	if x, ok := _PasswordResultTypeValue[name]; ok {
		return x, nil
	}
	return PasswordResultType(0), fmt.Errorf("%s is not a valid PasswordResultType", name)
}

// MarshalText implements the text marshaller method
func (x *PasswordResultType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *PasswordResultType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePasswordResultType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
