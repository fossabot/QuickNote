package sql

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type SQL struct {
	db *gorm.DB
}

func New(dialector gorm.Dialector) (*SQL, error) {
	db, err := gorm.Open(dialector, GetConfig())
	if err != nil {
		return nil, err
	}
	return &SQL{
		db: db,
	}, db.Error
}

func (d *SQL) Get() *gorm.DB {
	return d.db
}

func (d *SQL) Uninitialize() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func GetConfig() *gorm.Config {
	return &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	}
}
