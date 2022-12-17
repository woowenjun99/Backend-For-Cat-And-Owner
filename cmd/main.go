package main

import (
	// "context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	// "github.com/wenjun99/be/db"
	// "github.com/wenjun99/be/internal/handlers"
)

func SetConfig() {
	// Getting the environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalln("No .env file found")
	}



	// handlers.SetConfig(client)
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
