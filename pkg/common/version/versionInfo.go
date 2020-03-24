//
// The version command can be just added to your Cobra root command.
// At build time, the variables name, version, Commit, and BuildTags can be passed as build flags as shown in the following example:
//
//  go build -X "github.com/ver13/ava/pkg/common/version.name=Golang Microservices Framework" \
//           -X "github.com/ver13/ava/pkg/common/version.ServerName=avaServer" \
//           -X "github.com/ver13/ava/pkg/common/version.ClientName=avaCli" \
//           -X "github.com/ver13/ava/pkg/common/version.version=v1.0.0" \
//           -X "github.com/ver13/ava/pkg/common/version.Commit=f0f7b7dab7e36c20b757cebce0e8f4fc5b95de60" \
//           -X "github.com/ver13/ava/pkg/common/version.BuildTags=linux darwin amd64"
package version

import (
	"fmt"
	"runtime"
	"sync"

	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

var (
	l *VersionInfo

	once sync.Once
)

var (
	// application's name
	Name = ""
	// server binary name
	ServerName = "<appd>"
	// client binary name
	ClientName = "<appcli>"
	// application's version
	Version string
	Release = ""
	// commit
	Commit = ""
	// hash of the go.sum file
	GoSumHash = ""
	// build tags
	BuildTags   = ""
	BuildNumber = ""
	BuildDate   = ""
	BuildHash   = ""

	SemanticVersionStr = ""
)

type VersionInfo struct {
	Name       string `json:"name"`
	ServerName string `json:"server_name"`
	ClientName string `json:"client_name"`
	GitCommit  string `json:"commit"`

	SemanticVersion *SemanticVersion `json:"semantic_version"`

	GoVersion string `json:"go"`
}

func init() {
	once.Do(func() {
		semanticVersion, err := Parse(SemanticVersionStr)
		if err != nil {
			panic(err)
		}
		l = &VersionInfo{
			SemanticVersion: semanticVersion,
			Name:            Name,
			ServerName:      ServerName,
			ClientName:      ClientName,
			GitCommit:       Commit,
			GoVersion:       fmt.Sprintf("go version %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
		}
	})
}

func GetInstance() *VersionInfo {
	return l
}

func (v *VersionInfo) GetSemanticVersion() *SemanticVersion {
	return v.SemanticVersion
}

func (v *VersionInfo) GetName() string {
	return v.Name
}

func (v *VersionInfo) GetServerName() string {
	return v.ServerName
}

func (v *VersionInfo) GetClientName() string {
	return v.ClientName
}

func (v VersionInfo) GetGitCommit() string {
	return v.GitCommit
}

func (v *VersionInfo) GetGoVersion() string {
	return v.GoVersion
}

func (v *VersionInfo) String() string {
	s := serializerAVA.GetSerializer(serializerAVA.SerializerTypeToml)
	if s == nil {
		panic(fmt.Sprintf("Not exit serializer: %s", s.String()))
	}
	if a, err := s.Serializer(v); err != nil {
		return ""
	} else {
		return string(a)
	}
}
