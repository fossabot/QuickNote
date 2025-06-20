package orm

import (
	"gorm.io/gorm"
)

type Database interface {
	Get() *gorm.DB
	Uninitialize() error
}
