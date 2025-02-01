package config

import (
	"app/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load(".env")
	dbType := os.Getenv("DB_TYPE")

	var dsn string
	var errDB error

	if dbType == "mysql" {
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)
		DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbType == "sqlite" {
		dbName := os.Getenv("DB_NAME")
		dsn = fmt.Sprintf("%s.db", dbName)
		DB, errDB = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		panic("Unsupported DB_TYPE")
	}

	if errDB != nil {
		panic("Failed to Connect Database")
	}

	InitMigrate()

	fmt.Println("Connected to Database")
}

func InitMigrate() {
	err := DB.AutoMigrate(&model.Bot{}, &model.World{}, &model.Word{})
	if err != nil {
		return
	}
}
