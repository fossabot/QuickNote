package sql

import (
	"time"

	"github.com/Sn0wo2/QuickNote/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

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

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

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
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		FullSaveAssociations:   false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(zap.NewStdLog(log.Instance), logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		}),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
