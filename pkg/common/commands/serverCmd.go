package commands

import (
	"errors"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "serverName",
	Short:   "Start serverName.",
	Long:    "Start ServerName.",
	Example: `serverName start`,
	RunE:    serverStartCmdF,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func serverStartCmdF(command *cobra.Command, args []string) error {
	if _, err := InitCommandContext(command); err != nil {
		return err.Error()
	}

	if len(args) < 1 {
		return errors.New("Expected at least one argument. See help text for details.")
	}

	panic("Not implemented.")

	return nil
}
