package sql

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type SQL struct {
	db *gorm.DB
}

func New(dialector gorm.Dialector, config ...*gorm.Config) (*SQL, error) {
	var cfg *gorm.Config
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = GetConfig()
	}
	db, err := gorm.Open(dialector, cfg)
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
