package error

import (
	"encoding/json"
	"fmt"
)

// Error
type Error struct {
	Group    Group
	Subgroup Subgroup

	Err error

	Code    int
	Message string
	Details string

	Info *CallInfo
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) Println() string {
	return fmt.Sprintf("%s\n", e.String())
}

func (e *Error) String() string {
	return fmt.Sprintf("Group: %s, Subgroup: %s, Details: %s, Error: %v, Info: %v", e.Group, e.Subgroup, e.Details, e.Err, e.Info)
}

func (e *Error) ToJSON() string {
	var result []byte
	var err error

	result, err = json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}

	return string(result)
}
