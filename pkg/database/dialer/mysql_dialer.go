//go:build mysql

package dialer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mysqlDialector(dsn string) gorm.Dialector {
	return mysql.Open(dsn)
}
