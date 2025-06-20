//go:build postgres

package dialer

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresDialector(dsn string) gorm.Dialector {
	return postgres.Open(dsn)
}
