package version

import (
	"fmt"
	"strconv"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	. "github.com/ver13/ava/pkg/common/version/error"
)

// PrereleaseVersion represents a PreRelease SemanticVersion
type PrereleaseVersion struct {
	VersionStr string
	VersionNum uint64
	IsNum      bool
}

// NewPrereleaseVersion creates a new valid prerelease SemanticVersion
func NewPrereleaseVersion(s string) (*PrereleaseVersion, *errorAVA.Error) {
	if len(s) == 0 {
		return &PrereleaseVersion{}, PrereleaseIsEmpty(nil, "Prerelease is empty.")
	}
	v := PrereleaseVersion{}
	if containsOnly(s, numbers) {
		if hasLeadingZeroes(s) {
			return &PrereleaseVersion{}, PreReleaseWrong(nil, fmt.Sprintf("Numeric PreRelease SemanticVersion must not contain leading zeroes %q.", s))
		}
		num, err := strconv.ParseUint(s, 10, 64)

		// Might never be hit, but just in case
		if err != nil {
			return &PrereleaseVersion{}, PreReleaseWrong(err, err.Error())
		}
		v.VersionNum = num
		v.IsNum = true
	} else if containsOnly(s, alphanum) {
		v.VersionStr = s
		v.IsNum = false
	} else {
		return &PrereleaseVersion{}, PreReleaseWrong(nil, fmt.Sprintf("Invalid character(s) found in prerelease %q.", s))
	}
	return &v, nil
}

// Validate validates v and returns error in case
func (pre *PrereleaseVersion) Validate() *errorAVA.Error {
	if !pre.IsNum { //Numeric prerelease versions already uint64
		if len(pre.VersionStr) == 0 {
			return PreReleaseWrong(nil, fmt.Sprintf("Prerelease can not be empty %q", pre.VersionStr))
		}
		if !containsOnly(pre.VersionStr, alphanum) {
			return PreReleaseWrong(nil, fmt.Sprintf("Invalid character(s) found in prerelease %q", pre.VersionStr))
		}
	}

	return nil
}

// IsNumeric checks if prerelease-*SemanticVersion is numeric
func (pre *PrereleaseVersion) IsNumeric() bool {
	return pre.IsNum
}

// Compare compares two PreRelease Versions v and o:
// -1 == v is less than o
// 0 == v is equal to o
// 1 == v is greater than o
func (pre *PrereleaseVersion) Compare(o *PrereleaseVersion) int {
	if pre.IsNum && !o.IsNum {
		return -1
	} else if !pre.IsNum && o.IsNum {
		return 1
	} else if pre.IsNum && o.IsNum {
		if pre.VersionNum == o.VersionNum {
			return 0
		} else if pre.VersionNum > o.VersionNum {
			return 1
		} else {
			return -1
		}
	} else { // both are Alphas
		if pre.VersionStr == o.VersionStr {
			return 0
		} else if pre.VersionStr > o.VersionStr {
			return 1
		} else {
			return -1
		}
	}
}

// PreRelease *SemanticVersion to string
func (pre *PrereleaseVersion) String() string {
	if pre.IsNum {
		return strconv.FormatUint(pre.VersionNum, 10)
	}
	return pre.VersionStr
}
