package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (app *AppConfig) Test(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	if err := app.DB.Connect(app.ctx); err != nil {
		log.Fatalln("Error connecting to the db")
	}

	defer app.DB.Disconnect(app.ctx)

	databases, err := app.DB.ListDatabaseNames(app.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	response["databases"] = databases
	json.NewEncoder(w).Encode(response)
}
