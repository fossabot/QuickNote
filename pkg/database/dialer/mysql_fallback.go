//go:build !mysql
// +build !mysql

package dialer

import "gorm.io/gorm"

func mysqlDialector(dsn string) gorm.Dialector {
	panic("mysql driver not enabled: build with '-tags mysql'")
}
