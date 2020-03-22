package version

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type SemanticVersionI interface {
	String() string
	Equals(o *SemanticVersion) bool
	EQ(o *SemanticVersion) bool
	NE(o *SemanticVersion) bool
	GT(o *SemanticVersion) bool
	GTE(o *SemanticVersion) bool
	GE(o *SemanticVersion) bool
	LT(o *SemanticVersion) bool
	LTE(o *SemanticVersion) bool
	LE(o *SemanticVersion) bool
	Compare(o *SemanticVersion) int
	IncrementPatch() *errorAVA.Error
	IncrementMinor() *errorAVA.Error
	IncrementMajor() *errorAVA.Error
	Validate() *errorAVA.Error
}
