package version

import (
	"fmt"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	. "github.com/ver13/ava/pkg/common/version/error"
)

type BuildVersion struct {
	BuildTags   string
	BuildNumber string
	BuildDate   time.Time
	BuildHash   string
}

// Validate validates v and returns error in case
func (b *BuildVersion) Validate() *errorAVA.Error {
	if len(b.BuildNumber) == 0 {
		return BuildVersionIsEmpty(nil, fmt.Errorf("Build meta data can not be empty %q,", b))
	}
	if !containsOnly(b.BuildNumber, alphanum) {
		return BuildVersionWrong(nil, fmt.Errorf("Invalid character(s) found in build meta data %q.", b))
	}
	return nil
}

func (b *BuildVersion) String() string {
	return fmt.Sprintf("BuildNumber: [%s]\nBuildTags: [%s]\nBuildDate: [%s]\nBuildHash: [%s].", b.BuildNumber, b.BuildTags, b.BuildDate, b.BuildHash)
}
