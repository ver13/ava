package commands

import (
	"github.com/spf13/cobra"
)

var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Management of slash commands",
}

func init() {
	commandCmd.AddCommand()
	RootCmd.AddCommand(commandCmd)
}
