package version

import (
	"github.com/coreos/go-semver/semver"
)

type VersionInfoI interface {
	GetSemanticVersion() *semver.Version
	GetName() string
	GetServerName() string
	GetClientName() string
	GetGitCommit() string
	GetGoVersion() string
	String() string
}
