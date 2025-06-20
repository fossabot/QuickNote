//go:build sqlserver

package dialer

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func sqlserverDialector(dsn string) gorm.Dialector {
	return sqlserver.Open(dsn)
}
