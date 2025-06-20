package dialer

import (
	"strings"

	"gorm.io/gorm"
)

func New(name, dsn string) gorm.Dialector {
	switch strings.ToLower(name) {
	case "postgres", "pg", "cockroach", "crdb", "alloydb":
		return postgresDialector(dsn)
	case "sqlserver", "mssql":
		return sqlserverDialector(dsn)
	case "mysql", "mariadb", "tidb", "aurora":
		return mysqlDialector(dsn)
	// fallback or sqlite
	default:
		return sqliteDialector(dsn)
	}
}
