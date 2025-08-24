package config

import (
	"fmt"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(cfg *Config) (*gorm.DB, error) {
	databaseConfig := cfg.Database
	host := databaseConfig.Host
	port := databaseConfig.Port
	user := databaseConfig.User
	password := databaseConfig.Password
	dbname := databaseConfig.Name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	environment := cfg.Environment
	if environment == "development" {
		gormAutoMigrate(db)
	}

	return db, nil
}

func gormAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	}
}
