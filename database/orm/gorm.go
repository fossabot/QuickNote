package orm

import (
	"github.com/Sn0wo2/QuickNote/database/dialer"
	"github.com/Sn0wo2/QuickNote/database/sql"
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

func New(typ, url string) (Database, error) {
	return sql.New(dialer.New(typ, url))
}
