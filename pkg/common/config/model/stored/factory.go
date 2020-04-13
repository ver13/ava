package stored

import (
	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	cryptoAVA "github.com/ver13/ava/pkg/common/crypto"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

func Factory(t DialectType, name string, host string, port uint64, user string, password string, ssl SSLType, dbName string, debug bool, migrate bool) (*DbSQL, *errorAVA.Error) {
	switch t {
	case DialectTypeSQLServer:
		return newDbSQLServer(t, name, host, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	case DialectTypePostgreSQL:
		return newDbPostgres(t, name, host, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	case DialectTypeSqlite3:
		return newDbSqlite3(t, name, host, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	case DialectTypeMySQL:
		return newDbMySQL(t, name, host, port, user, cryptoAVA.NewPassword(password), ssl, dbName, debug, migrate)
	default:
		return nil, errorConfigAVA.DialectTypeUnknown(nil, t)
	}
}
