package main

import (
	"log"
	"os"

	"github.com/baguseka01/golang-jwt-authentication/database"
	"github.com/baguseka01/golang-jwt-authentication/routes"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func main() {
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	// Initialize Database
	database.Connect()
	database.Migrate()

	// Initialize Router
	router := routes.Router()
	router.Run(":" + appConfig.AppPort)
}
