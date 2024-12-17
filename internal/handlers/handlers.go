package handlers

import (
	"dms-backend/internal/db"
	"dms-backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	// var dbInstance *db.Database.ConnectDB
	// log.Fatal(dbInstance)
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Error decoding input: %v", err)
		return
	}
	err = db.CreateUser(&newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("Error inserting user: %v", err)
		return
	}

	// Respond with the created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	var input models.Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Error decoding input: %v", err)
		return
	}
	firstname, lastname, err := db.GetUser(&input)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		log.Printf("Error getting user: %v", err)
		return
	}
	response := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{
		FirstName: firstname,
		LastName:  lastname,
	}

	w.WriteHeader(http.StatusOK)
	// Encode the response struct into JSON
	json.NewEncoder(w).Encode(response)

}
