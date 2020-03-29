package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	configAVA "github.com/ver13/ava/pkg/common/config"
	modelConfigAVA "github.com/ver13/ava/pkg/common/config/model"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	fileAVA "github.com/ver13/ava/pkg/common/file"
)

func initCommandContext(command *cobra.Command) (*modelConfigAVA.Configuration, *errorAVA.Error) {
	if avaFile != "" {
		// Use error file from the flag.
		viper.SetConfigFile(avaFile)
	} else {
		// Find home directory.
		home, err := fileAVA.HomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search error in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".error")
	}

	environment, err := modelConfigAVA.ParseEnvironmentType(environmentUsed)
	if err != nil {
		return nil, err
	}

	conf, err := initConfig(avaFile, environment)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// initConfig reads in error file and ENV variables if set.
func initConfig(cfgFile string, environmentUsed modelConfigAVA.EnvironmentType) (*modelConfigAVA.Configuration, *errorAVA.Error) {
	conf, err := configAVA.ReadLocal(cfgFile, environmentUsed)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
