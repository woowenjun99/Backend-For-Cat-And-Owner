package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wenjun99/be/internal/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetConfig() {
	// Getting the environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalln("No .env file found")
	}

	connection_string := os.Getenv("DATABASE_URL")
	if connection_string == "" {
		log.Fatalln("There is no connection string")
	}

	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to the db", err.Error())
	}

	handlers.SetConfig(db)
}

func main() {
	SetConfig()
	srv := http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: Routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("Error starting up the server", err.Error())
	}
}
