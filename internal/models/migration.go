package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	// Getting the environment variables
	connection_string := os.Getenv("DATABASE_URL")
	if connection_string == "" {
		log.Fatalln("There is no connection string")
	}

	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalln("Error connecting to the db", err.Error())
	}

	return db
}

func Migration() {
	db := initDb()
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalln("Error migrating the user model: ", err.Error())
	}

	if err := db.AutoMigrate(&Event{}); err != nil {
		log.Fatalln("Error migrating the user model: ", err.Error())
	}

	log.Println("Successfully migrated all models")
}
