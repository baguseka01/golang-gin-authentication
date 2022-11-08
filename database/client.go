package database

import (
	"fmt"
	"log"
	"os"

	"github.com/baguseka01/golang-jwt-authentication/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbError error

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	DB, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
