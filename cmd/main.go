package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/wenjun99/be/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetConfig() {
	// Getting the environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalln("No .env file found")
	}

	connection_string := os.Getenv("DATABASE_URL")
	if connection_string == "" {
		log.Fatalln("The connection string is empty")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
	if err != nil {
		log.Fatalln("Error connecting to the database")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	handlers.App = handlers.SetConfig(client, ctx)
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
