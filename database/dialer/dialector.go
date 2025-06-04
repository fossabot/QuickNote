package dialer

import (
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func New(driver, url string) gorm.Dialector {
	switch strings.ToLower(driver) {
	case "postgres":
		return postgres.Open(url)
	case "sqlite":
		return sqlite.Open(url)
	case "sqlserver":
		return sqlserver.Open(url)
	// MySQL and fallback
	default:
		return mysql.Open(url)
	}
}
