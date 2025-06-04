package orm

import (
	"github.com/Sn0wo2/QuickNote/database/dialer"
	"github.com/Sn0wo2/QuickNote/database/sql"
	"gorm.io/gorm"
)

var Instance Database

func Init(typ, url string) error {
	db, err := New(typ, url)
	if err != nil {
		return err
	}
	Instance = db
	return Instance.Get().Error
}

type Database interface {
	Get() *gorm.DB
	Uninitialize() error
}

func New(typ, url string) (Database, error) {
	return sql.New(dialer.New(typ, url))
}
