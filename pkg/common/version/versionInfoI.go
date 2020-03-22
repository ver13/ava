package version

type VersionInfoI interface {
	GetSemanticVersion() *SemanticVersion
	GetName() string
	GetServerName() string
	GetClientName() string
	GetGitCommit() string
	GetGoVersion() string
	String() string
}
