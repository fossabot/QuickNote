//go:build !sqlserver
// +build !sqlserver

package dialer

import "gorm.io/gorm"

func sqlserverDialector(dsn string) gorm.Dialector {
	panic("sqlserver driver not enabled: build with '-tags sqlserver'")
}
