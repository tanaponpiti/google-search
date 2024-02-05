package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var GormDB *gorm.DB

func InitDatabase(connectionString string) error {
	var err error
	GormDB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := GormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
