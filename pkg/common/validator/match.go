package validator

import (
	"fmt"
	"regexp"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorValidatorAVA "github.com/ver13/ava/pkg/common/validator/error"
)

func match(text string, regex *regexp.Regexp) ([]string, *errorAVA.Error) {
	parsed := regex.FindAllString(text, -1)
	if parsed == nil {
		return nil, errorValidatorAVA.NotFoundData(nil, fmt.Sprintf("%s: %v", text, regex))
	}

	return parsed, nil
}
