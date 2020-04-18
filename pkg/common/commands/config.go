package commands

import (
	"errors"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration",
}

var validateConfigCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate config file",
	Long:  "If the config file is valid, this command will output a success message and have a zero exit code. If it is invalid, this command will output an error and have a non-zero exit code.",
	RunE:  configStartCmdF,
}

var configSubpathCmd = &cobra.Command{
	Use:   "subpath",
	Short: "Update client asset loading to use the configured subpath",
	Long:  "Update the hard-coded production client asset paths to take into account Gmf running on a subpath.",
	Example: `  config subpath
  config subpath --path /gmf
  config subpath --path /`,
	RunE: configStartCmdF,
}

func init() {
	configSubpathCmd.Flags().String("path", "", "Optional subpath; defaults to value in SiteURL")

	configCmd.AddCommand(
		validateConfigCmd,
		configSubpathCmd,
	)
	RootCmd.AddCommand(configCmd)
}

func configStartCmdF(command *cobra.Command, args []string) error {
	if _, err := InitCommandContext(command); err != nil {
		return err.Error()
	}

	if len(args) < 1 {
		return errors.New("Expected at least one argument. See help text for details.")
	}

	return nil
}
