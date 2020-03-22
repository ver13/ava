package validator

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CheckNewPassword Run some basic checks on new password strings, based on given options
// This routine requires at least 4 (four) characters
// Example requiring only basic minimum length: CheckNewPassword("lalala", "lalala", 10, CheckNewPasswordComplexityLowest)
// Example requiring number and symbol: CheckNewPassword("lalala", "lalala", 10, CheckNewPasswordComplexityRequireNumber|CheckNewPasswordComplexityRequireSymbol)
func CheckNewPassword(password, passwordConfirmation string, minimumlength uint, flagComplexity PasswordComplexityRulesType) PasswordResultType {
	const minPasswordLengthDefault = 4

	if minimumlength < minPasswordLengthDefault {
		minimumlength = 4
	}

	if utf8.RuneCountInString(strings.TrimSpace(password)) < int(minimumlength) {
		return PasswordResultTooShort
	}

	if password != passwordConfirmation {
		return PasswordResultDivergent
	}

	letterFound := false
	numberFound := false
	symbolFound := false
	spaceFound := false
	upperCaseFound := false

	if flagComplexity&PasswordComplexityLowest != PasswordComplexityLowest {
		for _, r := range password {
			if unicode.IsLetter(r) {
				letterFound = true

				if unicode.IsUpper(r) {
					upperCaseFound = true
				}
			}

			if unicode.IsNumber(r) {
				numberFound = true
			}

			if RuneHasSymbol(r) {
				symbolFound = true
			}

			if r == ' ' {
				spaceFound = true
			}
		}
	}

	if flagComplexity&PasswordComplexityRequireLetter == PasswordComplexityRequireLetter {
		if !letterFound {
			return PasswordResultTooSimple
		}

		// Only checks uppercase if letter is required
		if flagComplexity&PasswordComplexityRequireUpperCase == PasswordComplexityRequireUpperCase {
			if !upperCaseFound {
				return PasswordResultTooSimple
			}
		}
	}

	if flagComplexity&PasswordComplexityRequireNumber == PasswordComplexityRequireNumber {
		if !numberFound {
			return PasswordResultTooSimple
		}
	}

	if flagComplexity&PasswordComplexityRequireSymbol == PasswordComplexityRequireSymbol {
		if !symbolFound {
			return PasswordResultTooSimple
		}
	}

	if flagComplexity&PasswordComplexityRequireSpace == PasswordComplexityRequireSpace {
		if !spaceFound {
			return PasswordResultTooSimple
		}
	}

	return PasswordResultOK
}

// RuneHasSymbol returns true if the given rune contains a symbol
func RuneHasSymbol(ru rune) bool {
	allowedSymbols := "!\"#$%&'()*+´-./:;<=>?@[\\]^_`{|}~"

	for _, r := range allowedSymbols {
		if ru == r {
			return true
		}
	}

	return false
}
