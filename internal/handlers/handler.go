package handlers

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (app *AppConfig) Test() {
	if err := app.DB.Connect(app.ctx); err != nil {
		log.Fatalln("Error connecting to the db")
	}

	defer app.DB.Disconnect(app.ctx)

	databases, err := app.DB.ListDatabaseNames(app.ctx, bson.M{})
	if err != nil {
        log.Fatal(err)
    }

	log.Println(databases)
}
