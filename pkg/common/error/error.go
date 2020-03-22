package error

import (
	"encoding/json"
	"errors"
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

func (e *Error) Error() error {
	return errors.New(e.String())
}

func (e *Error) Println() string {
	return fmt.Sprintf("%s\n", e.String())
}

func (e *Error) String() string {
	return fmt.Sprintf("Group: %s, Subgroup: %s, Details: %s, Error: %v, Info: %v", e.Group, e.Subgroup, e.Details, e.Err, e.Info)
}

func (e *Error) ToJSON() (string, *Error) {
	var result []byte
	var err error

	result, err = json.MarshalIndent(e, "", "    ")
	if err != nil {
		err := Error{
			Group:    GroupGeneral,
			Subgroup: SubgroupSerializer,
			Details:  fmt.Sprintf("%v.", "Failed to generate JSON."),
			Err:      err,
			Code:     serializerJSONCode,
			Message:  statusTextFunc(serializerJSONCode),
			Info:     RetrieveCallInfo(),
		}

		return "", &err
	}

	return string(result), nil
}
