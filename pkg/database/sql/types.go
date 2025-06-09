package sql

import (
	"gorm.io/gorm"
)

type SQL struct {
	db *gorm.DB
}
