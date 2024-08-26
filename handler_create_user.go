package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Name string `json:"name"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode the request body into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	currentTime := time.Now()

	response := map[string]string{
		"id":         id.String(),
		"created_at": currentTime.Format(time.RFC3339),
		"updated_at": currentTime.Format(time.RFC3339),
		"name":       user.Name,
	}

	respondWithJSON(w, http.StatusOK, response)
}
