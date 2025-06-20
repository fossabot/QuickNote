//go:build !postgres
// +build !postgres

package dialer

import "gorm.io/gorm"

func postgresDialector(dsn string) gorm.Dialector {
	panic("postgres driver not enabled: build with '-tags postgres'")
}
