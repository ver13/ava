package stored

import (
	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	cryptoAVA "github.com/ver13/ava/pkg/common/crypto"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

const (
	PortPostgreSQLDefault = 5432
	PortMySQLDefault      = 3306
	PortSQLServerDefault  = 3306
	PortSqlite3Default    = 11111
)

type DbSQL struct {
	Dialect     DialectType
	URL         string
	Name        string
	User        string
	Password    *cryptoAVA.Password
	Port        uint64
	SSL         string
	DBName      string
	Debug       bool
	AutoMigrate bool
}

func newDbPostgres(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DbSQL, *errorAVA.Error) {
	if dialect != DialectTypePostgreSQL {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = PortPostgreSQLDefault
	}

	return &DbSQL{
		Dialect:     DialectTypePostgreSQL,
		URL:         url,
		Name:        name,
		User:        user,
		Password:    password,
		Port:        port,
		SSL:         ssl,
		DBName:      dbName,
		Debug:       debug,
		AutoMigrate: migrate,
	}, nil
}

func newDbMySQL(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DbSQL, *errorAVA.Error) {
	if dialect != DialectTypeMySQL {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = PortMySQLDefault
	}

	return &DbSQL{
		Dialect:     DialectTypeMySQL,
		URL:         url,
		Name:        name,
		User:        user,
		Password:    password,
		Port:        port,
		SSL:         ssl,
		DBName:      dbName,
		Debug:       debug,
		AutoMigrate: migrate,
	}, nil
}

func newDbSqlite3(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DbSQL, *errorAVA.Error) {
	if dialect != DialectTypeSqlite3 {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = PortSqlite3Default
	}

	return &DbSQL{
		Dialect:     DialectTypeSqlite3,
		URL:         url,
		Name:        name,
		User:        user,
		Password:    password,
		Port:        port,
		SSL:         ssl,
		DBName:      dbName,
		Debug:       debug,
		AutoMigrate: migrate,
	}, nil
}

func newDbSQLServer(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DbSQL, *errorAVA.Error) {
	if dialect != DialectTypeSQLServer {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = PortSQLServerDefault
	}

	return &DbSQL{
		Dialect:     DialectTypeSQLServer,
		URL:         url,
		Name:        name,
		User:        user,
		Password:    password,
		Port:        port,
		SSL:         ssl,
		DBName:      dbName,
		Debug:       debug,
		AutoMigrate: migrate,
	}, nil
}
