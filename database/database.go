package database

import (
	"fmt"

	"github.com/balasl342/apm-server-elastic-go/config"
	"github.com/balasl342/apm-server-elastic-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.AppConfig.Database.DBHost, config.AppConfig.Database.DBUser, config.AppConfig.Database.DBPassword, config.AppConfig.Database.DBName, config.AppConfig.Database.DBPort)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.User{})
}
