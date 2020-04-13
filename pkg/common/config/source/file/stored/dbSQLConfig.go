package stored

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"

	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	storedModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/stored"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	fileAVA "github.com/ver13/ava/pkg/common/file"
	errorFileAVA "github.com/ver13/ava/pkg/common/file/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	validatorAVA "github.com/ver13/ava/pkg/common/validator"
)

type DbSQLConfig struct {
	Dialect     string `mapstructure:"dialect"`
	Host        string `mapstructure:"url"`
	Name        string `mapstructure:"name,omitempty"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Port        uint64 `mapstructure:"port"`
	SSL         bool   `mapstructure:"ssl,omitempty"`
	DBName      string `mapstructure:"dbName"`
	Debug       bool   `mapstructure:"debug,omitempty"`
	AutoMigrate bool   `mapstructure:"auto_migrate,omitempty"`
}

func NewDbSQL(dialect string, url string, name string, user string, password string, port uint64, ssl bool, dbName string, debug bool, autoMigrate bool) (*storedModelConfigAVA.DbSQL, *errorAVA.Error) {
	dbSQL := &DbSQLConfig{
		Dialect:     dialect,
		Host:        url,
		Name:        name,
		User:        user,
		Password:    password,
		Port:        port,
		SSL:         ssl,
		DBName:      dbName,
		Debug:       debug,
		AutoMigrate: autoMigrate,
	}
	return dbSQL.Parser()
}

func NewDbSQLDefault() (*storedModelConfigAVA.DbSQL, *errorAVA.Error) {
	dbSQL := &DbSQLConfig{
		Dialect:     "PostgreSQL",
		Host:        "localhost",
		Name:        "AVA",
		User:        "ava",
		Password:    "postgres",
		Port:        5432,
		SSL:         false,
		DBName:      "user_demo",
		Debug:       true,
		AutoMigrate: true,
	}
	return dbSQL.Parser()
}

func (database *DbSQLConfig) ReadLocal(fileName string) (*storedModelConfigAVA.DbSQL, *errorAVA.Error) {
	if fileName == "" {
		return nil, errorFileAVA.FileNotFount(nil, "Config file path is empty.")
	}

	_err := fileAVA.NewFile().FileExists(fileName)
	if _err != nil {
		return nil, _err
	}

	var _viper *viper.Viper
	_viper = viper.New()

	// enable VIPER to read environment Variables
	_viper.AutomaticEnv()

	ext := filepath.Ext(fileName)[1:]
	_viper.SetConfigType(ext)

	// Set the file name of the configurations file
	_viper.SetConfigFile(fileName)

	if err := _viper.ReadInConfig(); err != nil {
		return nil, errorFileAVA.ReadFile(err, fileName)
	}

	return database.Parser()
}

func (database *DbSQLConfig) Parser() (*storedModelConfigAVA.DbSQL, *errorAVA.Error) {

	dialect, errCheckDialect := _checkDialect(database.Dialect)
	if errCheckDialect != nil {
		return nil, errCheckDialect
	}

	ssl, errCheckSSL := _checkSSL(database.SSL)
	if errCheckDialect != nil {
		return nil, errCheckSSL
	}

	dbName, errDBName := _checkDBName(database.DBName)
	if errDBName != nil {
		return nil, errDBName
	}

	name, errName := _checkName(database.Name)
	if errName != nil {
		return nil, errName
	}

	host, errHost := _checkHost(database.Host)
	if errHost != nil {
		return nil, errHost
	}

	password, errPassword := _checkPassword(database.Password)
	if errPassword != nil {
		return nil, errPassword
	}

	user, errUser := _checkUser(database.User)
	if errUser != nil {
		return nil, errUser
	}

	port, errPort := _checkPort(database.Port)
	if errPort != nil {
		return nil, errPort
	}

	return storedModelConfigAVA.Factory(
		dialect,
		name,
		host,
		port,
		user,
		password,
		ssl,
		dbName,
		database.Debug,
		database.AutoMigrate)
}

func (database *DbSQLConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(database)
}

func _checkDialect(dialect string) (storedModelConfigAVA.DialectType, *errorAVA.Error) {
	dialectType, err := storedModelConfigAVA.ParseDialectType(dialect)
	if err != nil {
		return storedModelConfigAVA.DialectTypeUnknown, errorConfigAVA.InvalidConfig(nil, fmt.Sprintf("DbSQLConfig dialect incorrect. %s", dialect))
	}

	return dialectType, nil
}

func _checkSSL(ssl bool) (storedModelConfigAVA.SSLType, *errorAVA.Error) {
	var sslType storedModelConfigAVA.SSLType = storedModelConfigAVA.SSLTypeUnknown
	if ssl {
		sslType = storedModelConfigAVA.SSLTypeEnable
	} else {
		sslType = storedModelConfigAVA.SSLTypeDisable
	}
	return sslType, nil
}

func _checkDBName(dbName string) (string, *errorAVA.Error) {
	if dbName == "" {
		return "", errorConfigAVA.InvalidConfig(nil, "DBName is empty.")
	}

	return dbName, nil
}

func _checkName(name string) (string, *errorAVA.Error) {
	if name == "" {
		return "", errorConfigAVA.InvalidConfig(nil, "Name is empty.")
	}
	return name, nil
}

func _checkHost(host string) (string, *errorAVA.Error) {
	if host == "" {
		return "", errorConfigAVA.InvalidConfig(nil, "Host is empty.")
	}
	if err := validatorAVA.CheckURL(host); err != nil {
		return "", err
	}
	return host, nil
}

func _checkPassword(password string) (string, *errorAVA.Error) {
	if password == "" {
		return "", errorConfigAVA.InvalidConfig(nil, "Password is empty.")
	}
	if _, err := validatorAVA.GetInstance().CheckPassword(password); err != nil {
		return "", err
	}
	return password, nil
}

func _checkUser(user string) (string, *errorAVA.Error) {
	if user == "" {
		return "", errorConfigAVA.InvalidConfig(nil, "User is empty.")
	}
	if _, err := validatorAVA.GetInstance().CheckUsername(user); err != nil {
		return "", err
	}
	return user, nil
}

func _checkPort(port uint64) (uint64, *errorAVA.Error) {
	if port == 0 {
		return 0, errorConfigAVA.InvalidConfig(nil, fmt.Sprintf("Port is empty or 0. %d", port))
	}
	if _, err := validatorAVA.GetInstance().CheckNotKnownPorts(strconv.FormatUint(port, 10)); err != nil {
		return 0, err
	}
	return port, nil
}
