package error

const (
	statusWildcardUnknown                   = 1
	statusBuildVersionIsEmpty               = 2
	statusBuildVersionIsWrong               = 3
	statusIncrementMajorError               = 4
	statusIncrementMinorError               = 5
	statusIncrementPatchError               = 6
	statusPrereleaseIsEmpty                 = 7
	statusPreReleaseWrong                   = 8
	statusSemanticVersionParseError         = 9
	statusSemanticVersionParseTolerantError = 10
	statusSemanticVersionIsEmpty            = 11
	statusValidateError                     = 12
	statusMajorParseError                   = 13
	statusMinorParseError                   = 14
	statusPatchParseError                   = 15
	statusBuildVersionParseError            = 16
	statusBuildDateParseError               = 17
)

var statusText = map[int]string{
	statusWildcardUnknown:                   "Not valid wildcard type.",
	statusBuildVersionIsEmpty:               "Build version is empty.",
	statusBuildVersionIsWrong:               "Build version is wrong.",
	statusIncrementMajorError:               "Increment major error.",
	statusIncrementMinorError:               "Increment minor error.",
	statusIncrementPatchError:               "Increment patch error.",
	statusPrereleaseIsEmpty:                 "Prerelease is empty.",
	statusPreReleaseWrong:                   "PreRelease wrong.",
	statusSemanticVersionParseError:         "Semantic version parse error.",
	statusSemanticVersionParseTolerantError: "Semantic version parse tolerant error.",
	statusSemanticVersionIsEmpty:            "Semantic version is empty.",
	statusValidateError:                     "Validate error.",
	statusMajorParseError:                   "Major parse error,",
	statusMinorParseError:                   "Minor parse error,",
	statusPatchParseError:                   "Patch parse error.",
	statusBuildVersionParseError:            "Build version parse error.",
	statusBuildDateParseError:               "Build date parse error.",
}

// StatusText returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
