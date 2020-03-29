package commands

import (
	"github.com/spf13/cobra"

	versionAVA "github.com/ver13/ava/pkg/common/version"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Display version information.",
	Long:    "Display version information.",
	Example: `  avaServer version `,
	RunE:    versionCmdF,
}

func versionCmdF(command *cobra.Command, args []string) error {
	l := versionAVA.GetInstance()

	_, _ = CommandPrintln("Name: " + l.Name)
	_, _ = CommandPrintln("ServerName: " + l.ServerName)
	_, _ = CommandPrintln("ClientName: " + l.ClientName)
	
	_, _ = CommandPrintln("Version: " + l.SemanticVersion.String())

	b := len(l.SemanticVersion.Build)
	for i := 0; i < b; i++{
		_, _ = CommandPrintln("Build Number: " + l.SemanticVersion.Build[i].BuildNumber)
		_, _ = CommandPrintln("Build Date: " + l.SemanticVersion.Build[i].BuildDate.String())
		_, _ = CommandPrintln("Build Hash: " + l.SemanticVersion.Build[i].BuildHash)
		_, _ = CommandPrintln("Build tags: " + l.SemanticVersion.Build[i].BuildTags)
	}
	_, _ = CommandPrintln("Git commit: " + l.GitCommit)
	_, _ = CommandPrintln("GoVersion: " + l.GoVersion)
	return nil
}
