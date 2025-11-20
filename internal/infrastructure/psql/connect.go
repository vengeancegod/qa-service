package psql

import (
	"fmt"
	"log"

	"qa-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host(), cfg.User(), cfg.Password(), cfg.DBName(), cfg.Port(), cfg.SSLMode()) //"host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed connect to db: %v", err)
		return nil, err
	}

	log.Printf("successfull connect to db: %s", cfg.DBName())
	return db, nil
}
