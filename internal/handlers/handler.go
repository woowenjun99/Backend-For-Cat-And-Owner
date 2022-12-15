package handlers

import (
	"encoding/json"
	"net/http"
)

func (app *AppConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var form struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	// Decode the JSON
	json.NewDecoder(r.Body).Decode(&form)

}
