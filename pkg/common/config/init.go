package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/ver13/ava/pkg/common/config/model"
	fileSourceConfigAVA "github.com/ver13/ava/pkg/common/config/source/file"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	fileAVA "github.com/ver13/ava/pkg/common/file"
	errorFileAVA "github.com/ver13/ava/pkg/common/file/error"
)

var _viper *viper.Viper

func init() {
	_viper = viper.New()
}

// ReadLocal reads a local error file from the top-level directory fileName should be without the extension
func ReadLocal(fileName string, environmentUsed model.EnvironmentType) (*model.Configuration, *errorAVA.Error) {
	if fileName == "" {
		return nil, errorFileAVA.FileNotFount(nil, "Config file path is empty.")
	}

	_err := fileAVA.NewFile().FileExists(fileName)
	if _err != nil {
		return nil, _err
	}

	// enable VIPER to read environment Variables
	_viper.AutomaticEnv()

	ext := filepath.Ext(fileName)[1:]
	_viper.SetConfigType(ext)

	// Set the file name of the configurations file
	_viper.SetConfigFile(fileName)

	if err := _viper.ReadInConfig(); err != nil {
		return nil, errorFileAVA.ReadFile(err, fileName)
	}

	var configuration fileSourceConfigAVA.ConfigurationConfig
	errUnmarshal := _viper.Unmarshal(&configuration)
	if errUnmarshal != nil {
		return nil, errorFileAVA.UnmarshalFile(errUnmarshal, fmt.Sprintf("Unable to decode into struct, %v", configuration))
	}

	return configuration.Parser(environmentUsed)
}
