package version

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorVersionAVA "github.com/ver13/ava/pkg/common/version/error"
)

type comparator func(*SemanticVersion, *SemanticVersion) bool

var (
	compEQ comparator = func(v1 *SemanticVersion, v2 *SemanticVersion) bool {
		return v1.Compare(v2) == 0
	}
	compNE = func(v1 *SemanticVersion, v2 *SemanticVersion) bool {
		return v1.Compare(v2) != 0
	}
	compGT = func(v1 *SemanticVersion, v2 *SemanticVersion) bool {
		return v1.Compare(v2) == 1
	}
	compGE = func(v1 *SemanticVersion, v2 *SemanticVersion) bool {
		return v1.Compare(v2) >= 0
	}
	compLT = func(v1 *SemanticVersion, v2 *SemanticVersion) bool {
		return v1.Compare(v2) == -1
	}
	compLE = func(v1 *SemanticVersion, v2 *SemanticVersion) bool {
		return v1.Compare(v2) <= 0
	}
)

type semanticVersionRange struct {
	v *SemanticVersion
	c comparator
}

// rangeFunc creates a SemanticVersionRangeFunc from the given semanticVersionRange.
func (vr *semanticVersionRange) rangeFunc() SemanticVersionRangeFunc {
	return SemanticVersionRangeFunc(func(v *SemanticVersion) bool {
		return vr.c(v, vr.v)
	})
}

// SemanticVersionRangeFunc represents a range of versions.
// A SemanticVersionRangeFunc can be used to check if a SemanticVersion satisfies it:
//
//     range, err := semver.ParseRange(">1.0.0 <2.0.0")
//     range(semver.MustParse("1.1.1") // returns true
type SemanticVersionRangeFunc func(*SemanticVersion) bool

// OR combines the existing SemanticVersionRangeFunc with another SemanticVersionRangeFunc using logical OR.
func (rf SemanticVersionRangeFunc) OR(f SemanticVersionRangeFunc) SemanticVersionRangeFunc {
	return SemanticVersionRangeFunc(func(v *SemanticVersion) bool {
		return rf(v) || f(v)
	})
}

// AND combines the existing SemanticVersionRangeFunc with another SemanticVersionRangeFunc using logical AND.
func (rf SemanticVersionRangeFunc) AND(f SemanticVersionRangeFunc) SemanticVersionRangeFunc {
	return SemanticVersionRangeFunc(func(v *SemanticVersion) bool {
		return rf(v) && f(v)
	})
}

// ParseRange parses a range and returns a SemanticVersionRangeFunc.
// If the range could not be parsed an error is returned.
//
// Valid ranges are:
//   - "<1.0.0"
//   - "<=1.0.0"
//   - ">1.0.0"
//   - ">=1.0.0"
//   - "1.0.0", "=1.0.0", "==1.0.0"
//   - "!1.0.0", "!=1.0.0"
//
// A SemanticVersionRangeFunc can consist of multiple ranges separated by space:
// Ranges can be linked by logical AND:
//   - ">1.0.0 <2.0.0" would match between both ranges, so "1.1.1" and "1.8.7" but not "1.0.0" or "2.0.0"
//   - ">1.0.0 <3.0.0 !2.0.3-beta.2" would match every SemanticVersion between 1.0.0 and 3.0.0 except 2.0.3-beta.2
//
// Ranges can also be linked by logical OR:
//   - "<2.0.0 || >=3.0.0" would match "1.x.x" and "3.x.x" but not "2.x.x"
//
// AND has a higher precedence than OR. It's not possible to use brackets.
//
// Ranges can be combined by both AND and OR
//
//  - `>1.0.0 <2.0.0 || >3.0.0 !4.2.1` would match `1.2.3`, `1.9.9`, `3.1.1`, but not `4.2.1`, `2.1.1`
func ParseRange(s string) (SemanticVersionRangeFunc, *errorAVA.Error) {
	parts := splitAndTrim(s)
	orParts, err := splitORParts(parts)
	if err != nil {
		return nil, err
	}
	expandedParts, err := expandWildcardVersion(orParts)
	if err != nil {
		return nil, err
	}
	var orFn SemanticVersionRangeFunc = nil
	for _, p := range expandedParts {
		var andFn SemanticVersionRangeFunc = nil
		for _, ap := range p {
			opStr, vStr, err := splitComparatorVersion(ap)
			if err != nil {
				return nil, err
			}
			vr, err := buildVersionRange(opStr, vStr)
			if err != nil {
				return nil, err
			}
			rf := vr.rangeFunc()

			// Set function
			if andFn == nil {
				andFn = rf
			} else { // Combine with existing function
				andFn = andFn.AND(rf)
			}
		}
		if orFn == nil {
			orFn = andFn
		} else {
			orFn = orFn.OR(andFn)
		}

	}
	return orFn, nil
}

// splitORParts splits the already cleaned parts by '||'.
// Checks for invalid positions of the operator and returns an
// error if found.
func splitORParts(parts []string) ([][]string, *errorAVA.Error) {
	var ORparts [][]string
	last := 0
	for i, p := range parts {
		if p == "||" {
			if i == 0 {
				return nil, errorVersionAVA.SemanticVersionParseError(nil, fmt.Sprintf("First element in range is '||'"))
			}
			ORparts = append(ORparts, parts[last:i])
			last = i + 1
		}
	}
	if last == len(parts) {
		return nil, errorVersionAVA.SemanticVersionParseError(nil, fmt.Sprintf("Last element in range is '||'"))
	}
	ORparts = append(ORparts, parts[last:])
	return ORparts, nil
}

// buildVersionRange takes a slice of 2: operator and SemanticVersion
// and builds a semanticVersionRange, otherwise an error.
func buildVersionRange(opStr, vStr string) (*semanticVersionRange, *errorAVA.Error) {
	c := parseComparator(opStr)
	if c == nil {
		return nil, errorVersionAVA.BuildVersionWrong(nil, fmt.Sprintf("Could not parse comparator %q in %q", opStr, strings.Join([]string{opStr, vStr}, "")))
	}
	v, err := Parse(vStr)
	if err != nil {
		return nil, err
	}

	return &semanticVersionRange{
		v: v,
		c: c,
	}, nil

}

// inArray checks if a byte is contained in an array of bytes
func inArray(s byte, list []byte) bool {
	for _, el := range list {
		if el == s {
			return true
		}
	}
	return false
}

// splitAndTrim splits a range string by spaces and cleans whitespaces
func splitAndTrim(s string) (result []string) {
	last := 0
	var lastChar byte
	excludeFromSplit := []byte{'>', '<', '='}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && !inArray(lastChar, excludeFromSplit) {
			if last < i-1 {
				result = append(result, s[last:i])
			}
			last = i + 1
		} else if s[i] != ' ' {
			lastChar = s[i]
		}
	}
	if last < len(s)-1 {
		result = append(result, s[last:])
	}

	for i, v := range result {
		result[i] = strings.Replace(v, " ", "", -1)
	}

	// parts := strings.Split(s, " ")
	// for _, x := range parts {
	// 	if s := strings.TrimSpace(x); len(s) != 0 {
	// 		result = append(result, s)
	// 	}
	// }
	return
}

// splitComparatorVersion splits the comparator from the SemanticVersion.
// Input must be free of leading or trailing spaces.
func splitComparatorVersion(s string) (string, string, *errorAVA.Error) {
	i := strings.IndexFunc(s, unicode.IsDigit)
	if i == -1 {
		return "", "", errorVersionAVA.SemanticVersionParseError(nil, fmt.Sprintf("Could not get SemanticVersion from string: %q", s))
	}
	return strings.TrimSpace(s[0:i]), s[i:], nil
}

// getWildcardType will return the type of wildcard that the
// passed SemanticVersion contains
func getWildcardType(vStr string) WildcardType {
	parts := strings.Split(vStr, ".")
	nparts := len(parts)
	wildcard := parts[nparts-1]

	var possibleWildcardType = WildcardType(nparts)
	if wildcard == "x" {
		return possibleWildcardType
	}

	return WildcardTypeNone
}

// createVersionFromWildcard will convert a wildcard SemanticVersion
// into a regular SemanticVersion, replacing 'x's with '0's, handling
// special cases like '1.x.x' and '1.x'
func createVersionFromWildcard(vStr string) string {
	// handle 1.x.x
	vStr2 := strings.Replace(vStr, ".x.x", ".x", 1)
	vStr2 = strings.Replace(vStr2, ".x", ".0", 1)
	parts := strings.Split(vStr2, ".")

	// handle 1.x
	if len(parts) == 2 {
		return vStr2 + ".0"
	}

	return vStr2
}

// incrementMajorVersion will increment the major SemanticVersion
// of the passed SemanticVersion
func incrementMajorVersion(vStr string) (string, *errorAVA.Error) {
	parts := strings.Split(vStr, ".")
	i, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", errorVersionAVA.IncrementMajorError(err, vStr)
	}
	parts[0] = strconv.Itoa(i + 1)

	return strings.Join(parts, "."), nil
}

// incrementMajorVersion will increment the minor SemanticVersion
// of the passed SemanticVersion
func incrementMinorVersion(vStr string) (string, *errorAVA.Error) {
	parts := strings.Split(vStr, ".")
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", errorVersionAVA.IncrementMinorError(err, vStr)
	}
	parts[1] = strconv.Itoa(i + 1)

	return strings.Join(parts, "."), nil
}

// expandWildcardVersion will expand wildcards inside versions
// following these rules:
//
// * when dealing with patch wildcards:
// >= 1.2.x    will become    >= 1.2.0
// <= 1.2.x    will become    <  1.3.0
// >  1.2.x    will become    >= 1.3.0
// <  1.2.x    will become    <  1.2.0
// != 1.2.x    will become    <  1.2.0 >= 1.3.0
//
// * when dealing with minor wildcards:
// >= 1.x      will become    >= 1.0.0
// <= 1.x      will become    <  2.0.0
// >  1.x      will become    >= 2.0.0
// <  1.0      will become    <  1.0.0
// != 1.x      will become    <  1.0.0 >= 2.0.0
//
// * when dealing with wildcards without
// SemanticVersion operator:
// 1.2.x       will become    >= 1.2.0 < 1.3.0
// 1.x         will become    >= 1.0.0 < 2.0.0
func expandWildcardVersion(parts [][]string) ([][]string, *errorAVA.Error) {
	var expandedParts [][]string
	for _, p := range parts {
		var newParts []string
		for _, ap := range p {
			if strings.Contains(ap, "x") {
				opStr, vStr, err := splitComparatorVersion(ap)
				if err != nil {
					return nil, err
				}

				versionWildcardType := getWildcardType(vStr)
				flatVersion := createVersionFromWildcard(vStr)

				var resultOperator string
				var shouldIncrementVersion bool
				switch opStr {
				case ">":
					resultOperator = ">="
					shouldIncrementVersion = true
				case ">=":
					resultOperator = ">="
				case "<":
					resultOperator = "<"
				case "<=":
					resultOperator = "<"
					shouldIncrementVersion = true
				case "", "=", "==":
					newParts = append(newParts, ">="+flatVersion)
					resultOperator = "<"
					shouldIncrementVersion = true
				case "!=", "!":
					newParts = append(newParts, "<"+flatVersion)
					resultOperator = ">="
					shouldIncrementVersion = true
				}

				var resultVersion string
				if shouldIncrementVersion {
					switch versionWildcardType {
					case WildcardTypePatch:
						resultVersion, _ = incrementMinorVersion(flatVersion)
					case WildcardTypeMinor:
						resultVersion, _ = incrementMajorVersion(flatVersion)
					}
				} else {
					resultVersion = flatVersion
				}

				ap = resultOperator + resultVersion
			}
			newParts = append(newParts, ap)
		}
		expandedParts = append(expandedParts, newParts)
	}

	return expandedParts, nil
}

func parseComparator(s string) comparator {
	switch s {
	case "==":
		fallthrough
	case "":
		fallthrough
	case "=":
		return compEQ
	case ">":
		return compGT
	case ">=":
		return compGE
	case "<":
		return compLT
	case "<=":
		return compLE
	case "!":
		fallthrough
	case "!=":
		return compNE
	}

	return nil
}

// MustParseRange is like ParseRange but panics if the range cannot be parsed.
func MustParseRange(s string) (SemanticVersionRangeFunc, *errorAVA.Error) {
	r, err := ParseRange(s)
	if err != nil {
		return nil, err
	}
	return r, nil
}
