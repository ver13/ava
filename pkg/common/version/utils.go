package version

import (
	"strings"
)

func containsOnly(s string, set string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !strings.ContainsRune(set, r)
	}) == -1
}

func hasLeadingZeroes(s string) bool {
	return len(s) > 1 && s[0] == '0'
}
