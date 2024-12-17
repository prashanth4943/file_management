package handlers

import (
	"dms-backend/databuckets"
	"dms-backend/internal/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UUID  string `json:"uuid"`
		Email string `json:"email"`
	}
	var req Request

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body. Method: %s, URL: %s, Reason: %v", r.Method, r.URL.Path, err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if req.UUID == "" && req.Email == "" {
		http.Error(w, "At least one of 'uuid' or 'email' must be provided", http.StatusBadRequest)
		return
	}

	if req.Email != "" {
		err = databuckets.DeleteFromOracle(&req.Email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete object: %v", err), http.StatusInternalServerError)
			return
		}
	} else if req.UUID != "" {
		var email string
		query := "SELECT email FROM users WHERE uuid = ?"
		err = db.DBInstance.Conn.QueryRow(query, req.UUID).Scan(&email)
		if err != nil {
			http.Error(w, "error in query", http.StatusInternalServerError)
			log.Printf("error in query , reason: %s", err)
		}
		err = databuckets.DeleteFromOracle(&email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete object: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to delete object: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("File associated with %s deleted successfully.", req.Email)))
}
