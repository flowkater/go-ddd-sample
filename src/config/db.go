package config

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, err := conn.DB()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(50)
	sqlDb.SetMaxOpenConns(100)

	db = conn.Set("gorm:auto_update", false)

	return nil
}

func DB() *gorm.DB {
	return db
}

func DBWithContext(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
