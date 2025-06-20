//go:build !sqlite
// +build !sqlite

package dialer

import "gorm.io/gorm"

func sqliteDialector(dsn string) gorm.Dialector {
	panic("sqlite driver not enabled: build with '-tags sqlite'")
}
