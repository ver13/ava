//go:generate ava-enum -f=$GOFILE --marshal --lower

package stored

// DialectType x ENUM(
// PostgreSQL
// Sqlite3
// MySQL
// SQLServer
// Unknown
// )
type DialectType int32
