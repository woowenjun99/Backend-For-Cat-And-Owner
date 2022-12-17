package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/wenjun99/be/db"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Generic Settings
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]interface{})
	client := db.NewClient()
	defer client.Prisma.Disconnect()
	defer r.Body.Close()

	// Attempt to connect to the database
	if err := client.Prisma.Connect(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response["error"] = "Something went wrong"
		log.Fatalln("Error connecting to the database", err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return the search results
	result, err := client.User.FindMany().Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response["error"] = "Something went wrong"
		log.Println("Something went wrong when finding users: ", err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}
