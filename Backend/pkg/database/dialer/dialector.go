package dialer

import (
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func New(name, dsn string) gorm.Dialector {
	switch strings.ToLower(name) {
	case "postgres", "pg", "cockroach", "crdb", "alloydb":
		return postgres.Open(dsn)
	case "sqlite", "sqlite3":
		return sqlite.Open(dsn)
	case "sqlserver", "mssql":
		return sqlserver.Open(dsn)
	// fallback or mysql
	// "mysql", "mariadb", "tidb", "aurora"
	default:
		return mysql.Open(dsn)
	}
}
