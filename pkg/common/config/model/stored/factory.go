package stored

import (
	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	cryptoAVA "github.com/ver13/ava/pkg/common/crypto"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func Factory(t DialectType, name string, url string, port uint64, user string, password string, ssl string, dbName string, debug bool, migrate bool) (*DBSQL, *errorAVA.Error) {
	switch t {
	case DialectTypeSQLServer:
		return newDBSQLServer(t, name, url, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	case DialectTypePostgreSQL:
		return newDBPostgres(t, name, url, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	case DialectTypeSqlite3:
		return newDBSqlite3(t, name, url, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	case DialectTypeMySQL:
		return newDBMySQL(t, name, url, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	default:
		return nil, errorConfigAVA.DialectTypeUnknown(nil, t)
	}
}
