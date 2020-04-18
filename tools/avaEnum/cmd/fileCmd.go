package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ver13/ava/pkg/common/commands"
)

var fileCmd = &cobra.Command{
	Use:     "file",
	Short:   "The file(s) to generate enums.  Use more than one flag for more files.",
	Long:    "The file(s) to generate enums.  Use more than one flag for more files.",
	Example: `file exampleEnum.go`,
	RunE:    fileCmdF,
}

func init() {
	commands.RootCmd.AddCommand(fileCmd)
}

func fileCmdF(command *cobra.Command, args []string) error {
	if _, err := commands.InitCommandContext(command); err != nil {
		return err.Error()
	}

	if len(args) < 1 {
		return errors.New("Expected at least one argument. See help text for details.")
	}

	panic("Not implemented.")

	return nil
}
