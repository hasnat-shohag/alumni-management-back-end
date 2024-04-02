package connection

import (
	"alumni-management-server/pkg/config"
	"alumni-management-server/pkg/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// database connection
func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbIp, dbConfig.DbName)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	db = d
}

// create table
func migrate() {
	if err := db.Migrator().AutoMigrate(&models.UserDetail{}); err != nil {
		fmt.Println("Error migrating DB")
		panic(err)
	}
}

// function for getting db instance
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
