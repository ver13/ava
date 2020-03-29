package stored

import (
	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	cryptoAVA "github.com/ver13/ava/pkg/common/crypto"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

const (
	portPostgreSQLDefault = 5432
	portMySQLDefault      = 3306
	portSQLServerDefault  = 3306
	portSqlite3Default    = 11111
)

type DBSQL struct {
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

func newDBPostgres(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DBSQL, *errorAVA.Error) {
	if dialect != DialectTypePostgreSQL {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = portPostgreSQLDefault
	}

	return &DBSQL{
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

func newDBMySQL(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DBSQL, *errorAVA.Error) {
	if dialect != DialectTypeMySQL {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = portMySQLDefault
	}

	return &DBSQL{
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

func newDBSqlite3(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DBSQL, *errorAVA.Error) {
	if dialect != DialectTypeSqlite3 {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = portSqlite3Default
	}

	return &DBSQL{
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

func newDBSQLServer(dialect DialectType, name string, url string, port uint64, user string, password *cryptoAVA.Password, ssl string, dbName string, debug bool, migrate bool) (*DBSQL, *errorAVA.Error) {
	if dialect != DialectTypeSQLServer {
		return nil, errorConfigAVA.DialectIsWrong(nil, dialect)
	}

	if port <= 0 {
		port = portSQLServerDefault
	}

	return &DBSQL{
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
