package version

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorVersionAVA "github.com/ver13/ava/pkg/common/version/error"
)

const (
	numbers  string = "0123456789"
	alphas          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-"
	alphanum        = alphas + numbers
)

type SemanticVersion struct {
	Major uint64
	Minor uint64
	Patch uint64
	Pre   []PrereleaseVersion
	Build []BuildVersion
}

// SemanticVersion to string
func (v *SemanticVersion) String() string {
	b := make([]byte, 0, 5)
	b = strconv.AppendUint(b, v.Major, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, v.Minor, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, v.Patch, 10)

	if len(v.Pre) > 0 {
		b = append(b, '-')
		b = append(b, v.Pre[0].String()...)

		for _, pre := range v.Pre[1:] {
			b = append(b, '.')
			b = append(b, pre.String()...)
		}
	}

	if len(v.Build) > 0 {
		b = append(b, '+')
		b = append(b, v.Build[0].String()...)

		for _, build := range v.Build[1:] {
			b = append(b, '.')
			b = append(b, build.String()...)
		}
	}

	return string(b)
}

// Equals checks if v is equal to o.
func (v *SemanticVersion) Equals(o *SemanticVersion) bool {
	return v.Compare(o) == 0
}

// EQ checks if v is equal to o.
func (v *SemanticVersion) EQ(o *SemanticVersion) bool {
	return v.Compare(o) == 0
}

// NE checks if v is not equal to o.
func (v *SemanticVersion) NE(o *SemanticVersion) bool {
	return v.Compare(o) != 0
}

// GT checks if v is greater than o.
func (v *SemanticVersion) GT(o *SemanticVersion) bool {
	return v.Compare(o) == 1
}

// GTE checks if v is greater than or equal to o.
func (v *SemanticVersion) GTE(o *SemanticVersion) bool {
	return v.Compare(o) >= 0
}

// GE checks if v is greater than or equal to o.
func (v *SemanticVersion) GE(o *SemanticVersion) bool {
	return v.Compare(o) >= 0
}

// LT checks if v is less than o.
func (v *SemanticVersion) LT(o *SemanticVersion) bool {
	return v.Compare(o) == -1
}

// LTE checks if v is less than or equal to o.
func (v *SemanticVersion) LTE(o *SemanticVersion) bool {
	return v.Compare(o) <= 0
}

// LE checks if v is less than or equal to o.
func (v *SemanticVersion) LE(o *SemanticVersion) bool {
	return v.Compare(o) <= 0
}

// Compare compares Semantic Versions v to o:
// -1 == v is less than o
// 0 == v is equal to o
// 1 == v is greater than o
func (v *SemanticVersion) Compare(o *SemanticVersion) int {
	if v.Major != o.Major {
		if v.Major > o.Major {
			return 1
		}
		return -1
	}
	if v.Minor != o.Minor {
		if v.Minor > o.Minor {
			return 1
		}
		return -1
	}
	if v.Patch != o.Patch {
		if v.Patch > o.Patch {
			return 1
		}
		return -1
	}

	// Quick comparison if a Semantic Version has no prerelease versions
	if len(v.Pre) == 0 && len(o.Pre) == 0 {
		return 0
	} else if len(v.Pre) == 0 && len(o.Pre) > 0 {
		return 1
	} else if len(v.Pre) > 0 && len(o.Pre) == 0 {
		return -1
	}

	i := 0
	for ; i < len(v.Pre) && i < len(o.Pre); i++ {
		if comp := v.Pre[i].Compare(&o.Pre[i]); comp == 0 {
			continue
		} else if comp == 1 {
			return 1
		} else {
			return -1
		}
	}

	// If all pr versions are the equal but one has further PrereleaseVersion, this one greater
	if i == len(v.Pre) && i == len(o.Pre) {
		return 0
	} else if i == len(v.Pre) && i < len(o.Pre) {
		return -1
	} else {
		return 1
	}

}

// IncrementPatch increments the patch *SemanticVersion
func (v *SemanticVersion) IncrementPatch() *errorAVA.Error {
	if v.Major == 0 {
		return errorVersionAVA.IncrementPatchError(nil, fmt.Sprintf("Patch SemanticVersion can not be incremented for %q", v.String()))
	}
	v.Patch += 1
	return nil
}

// IncrementMinor increments the minor *SemanticVersion
func (v *SemanticVersion) IncrementMinor() *errorAVA.Error {
	if v.Major == 0 {
		return errorVersionAVA.IncrementMinorError(nil, fmt.Sprintf("Minor *SemanticVersion can not be incremented for %q", v.String()))
	}
	v.Minor += 1
	v.Patch = 0
	return nil
}

// IncrementMajor increments the major *SemanticVersion
func (v *SemanticVersion) IncrementMajor() *errorAVA.Error {
	if v.Major == 0 {
		return errorVersionAVA.IncrementMajorError(nil, fmt.Sprintf("Major *SemanticVersion can not be incremented for %q", v.String()))
	}
	v.Major += 1
	v.Minor = 0
	v.Patch = 0
	return nil
}

// Validate validates v and returns error in case
func (v *SemanticVersion) Validate() *errorAVA.Error {
	// Major, Minor, Patch already validated using uint64

	for _, pre := range v.Pre {
		if err := pre.Validate(); err != nil {
			return err
		}
	}

	for _, build := range v.Build {
		if err := build.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// New is an alias for Parse and returns a pointer, parses Semantic Version string and returns a validated SemanticVersion or error
func New(s string) (vp *SemanticVersion, err *errorAVA.Error) {
	v, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vp = v
	return
}

// Make is an alias for Parse, parses SemanticVersion string and returns a validated SemanticVersion or error
func Make(s string) (*SemanticVersion, *errorAVA.Error) {
	return Parse(s)
}

// ParseTolerant allows for certain SemanticVersion specifications that do not strictly adhere to SemanticVersion specs to be parsed by this library.
// It does so by normalizing versions before passing them to Parse().
// It currently trims spaces, removes a "v" prefix, adds a 0 patch number to versions with only major and minor components specified, and removes leading 0s.
func ParseTolerant(s string) (*SemanticVersion, *errorAVA.Error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "v")

	// Split into major.minor.(patch+pr+meta)
	parts := strings.SplitN(s, ".", 3)
	// Remove leading zeros.
	for i, p := range parts {
		if len(p) > 1 {
			parts[i] = strings.TrimPrefix(p, "0")
		}
	}
	// Fill up shortened versions.
	if len(parts) < 3 {
		if strings.ContainsAny(parts[len(parts)-1], "+-") {
			return &SemanticVersion{}, errorVersionAVA.SemanticVersionParseTolerantError(nil, "Short SemanticVersion cannot contain PreRelease/Build meta data")
		}
		for len(parts) < 3 {
			parts = append(parts, "0")
		}
	}
	s = strings.Join(parts, ".")

	return Parse(s)
}

// Parse parses SemanticVersion string and returns a validated Semantic Version or error
func Parse(s string) (*SemanticVersion, *errorAVA.Error) {
	if len(s) == 0 {
		return &SemanticVersion{}, errorVersionAVA.SemanticVersionIsEmpty(nil, "SemanticVersion string is empty.")
	}

	// Split into major.minor.(patch+pr+meta)
	parts := strings.SplitN(s, ".", 3)
	if len(parts) != 3 {
		return &SemanticVersion{}, errorVersionAVA.SemanticVersionParseError(nil, "No Major.Minor.Patch elements found")
	}

	// Major
	if !containsOnly(parts[0], numbers) {
		return &SemanticVersion{}, errorVersionAVA.MajorParseError(nil, fmt.Sprintf("Invalid character(s) found in major number %q", parts[0]))
	}
	if hasLeadingZeroes(parts[0]) {
		return &SemanticVersion{}, errorVersionAVA.MajorParseError(nil, fmt.Sprintf("Major number must not contain leading zeroes %q", parts[0]))
	}
	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return &SemanticVersion{}, errorVersionAVA.MajorParseError(err, err.Error())
	}

	// Minor
	if !containsOnly(parts[1], numbers) {
		return &SemanticVersion{}, errorVersionAVA.MinorParseError(nil, fmt.Sprintf("Invalid character(s) found in minor number %q", parts[1]))
	}
	if hasLeadingZeroes(parts[1]) {
		return &SemanticVersion{}, errorVersionAVA.MinorParseError(nil, fmt.Sprintf("Minor number must not contain leading zeroes %q", parts[1]))
	}
	minor, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return &SemanticVersion{}, errorVersionAVA.MinorParseError(err, err.Error())
	}

	v := &SemanticVersion{}
	v.Major = major
	v.Minor = minor

	var build, prerelease []string
	patchStr := parts[2]

	if buildIndex := strings.IndexRune(patchStr, '+'); buildIndex != -1 {
		build = strings.Split(patchStr[buildIndex+1:], ".")
		patchStr = patchStr[:buildIndex]
	}

	if preIndex := strings.IndexRune(patchStr, '-'); preIndex != -1 {
		prerelease = strings.Split(patchStr[preIndex+1:], ".")
		patchStr = patchStr[:preIndex]
	}

	if !containsOnly(patchStr, numbers) {
		return &SemanticVersion{}, errorVersionAVA.PatchParseError(nil, fmt.Sprintf("Invalid character(s) found in patch number %q", patchStr))
	}
	if hasLeadingZeroes(patchStr) {
		return &SemanticVersion{}, errorVersionAVA.PatchParseError(nil, fmt.Sprintf("Patch number must not contain leading zeroes %q", patchStr))
	}
	patch, err := strconv.ParseUint(patchStr, 10, 64)
	if err != nil {
		return &SemanticVersion{}, errorVersionAVA.PatchParseError(err, err.Error())
	}

	v.Patch = patch

	// Prerelease
	for _, prstr := range prerelease {
		parsedPR, err := NewPrereleaseVersion(prstr)
		if err != nil {
			return &SemanticVersion{}, err
		}
		v.Pre = append(v.Pre, *parsedPR)
	}

	// Build meta data
	for _, str := range build {
		parsedBuild, err := NewBuildVersion(str)
		if err != nil {
			return &SemanticVersion{}, err
		}
		v.Build = append(v.Build, *parsedBuild)
	}

	return v, nil
}

// NewBuildVersion creates a new valid build SemanticVersion
func NewBuildVersion(str string) (*BuildVersion, *errorAVA.Error) {
	if len(str) == 0 {
		return &BuildVersion{}, errorVersionAVA.BuildVersionIsEmpty(nil, "Build meta data is empty.")
	}
	if !containsOnly(str, alphanum) {
		return &BuildVersion{}, errorVersionAVA.BuildVersionParseError(nil, fmt.Sprintf("Invalid character(s) found in build meta data %q", str))
	}
	eparts := strings.Split(str, ".")

	t, err := time.Parse("", eparts[2])
	if err != nil {
		return &BuildVersion{}, errorVersionAVA.BuildDateParseError(nil, fmt.Sprintf("Invalid time format found in build meta data %s", eparts[2]))
	}
	return &BuildVersion{
		BuildTags:   eparts[0],
		BuildNumber: eparts[1],
		BuildDate:   t,
		BuildHash:   eparts[3],
	}, nil
}
