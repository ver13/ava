package stored

import (
	"fmt"
	"strconv"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	"github.com/ver13/ava/pkg/common/config/model/stored"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	validatorAVA "github.com/ver13/ava/pkg/common/validator"
)

type DatabaseSQL struct {
	Dialect     string `mapstructure:"dialect"`
	URL         string `mapstructure:"url"`
	Name        string `mapstructure:"name,omitempty"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Port        uint64 `mapstructure:"port"`
	SSL         bool   `mapstructure:"ssl,omitempty"`
	DBName      string `mapstructure:"dbName"`
	Debug       bool   `mapstructure:"debug,omitempty"`
	AutoMigrate bool   `mapstructure:"auto_migrate,omitempty"`
}

func (database *DatabaseSQL) Parser() (*stored.DBSQL, *errorAVA.Error) {

	dialect, err := stored.ParseDialectType(database.Dialect)
	if err != nil {
		return nil, errorConfigAVA.InvalidConfig(nil, fmt.Sprintf("DatabaseSQL dialect incorrect. %s", database.Dialect))
	}

	var ssl = "enable"
	if !database.SSL {
		ssl = "disable"
	}

	if database.DBName == "" {
		return nil, errorConfigAVA.InvalidConfig(nil, "DBName is empty.")
	}

	if database.Name == "" {
		return nil, errorConfigAVA.InvalidConfig(nil, "Name is empty.")
	}

	if database.URL == "" {
		return nil, errorConfigAVA.InvalidConfig(nil, "URL is empty.")
	}
	if err := validatorAVA.CheckURL(database.URL); err != nil {
		return nil, err
	}

	if database.Password == "" {
		return nil, errorConfigAVA.InvalidConfig(nil, "Password is empty.")
	}
	if _, err := validatorAVA.GetInstance().CheckPassword(database.Password); err != nil {
		return nil, err
	}

	if database.User == "" {
		return nil, errorConfigAVA.InvalidConfig(nil, "User is empty.")
	}
	if _, err := validatorAVA.GetInstance().CheckUsername(database.User); err != nil {
		return nil, err
	}

	if database.Port == 0 {
		return nil, errorConfigAVA.InvalidConfig(nil, fmt.Sprintf("Port is empty or 0. %s", database.Port))
	}
	if _, err := validatorAVA.GetInstance().CheckNotKnownPorts(strconv.FormatUint(database.Port, 10)); err != nil {
		return nil, err
	}

	return stored.Factory(
		dialect,
		database.Name,
		database.URL,
		database.Port,
		database.User,
		database.Password,
		ssl,
		database.DBName,
		database.Debug,
		database.AutoMigrate)
}

func (database *DatabaseSQL) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(database)
}
