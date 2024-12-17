package handlers

import (
	"dms-backend/databuckets"
	"dms-backend/internal/db"
	"encoding/json"
	"log"
	"net/http"
)

type DownloadResponse struct {
	Msg string `json:"url"`
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UUID string `json:"uuid"`
	}
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.UUID == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	var email string

	query := "SELECT email FROM users WHERE uuid = ?"
	err = db.DBInstance.Conn.QueryRow(query, request.UUID).Scan(&email)
	if err != nil {
		http.Error(w, "error in query", http.StatusInternalServerError)
		log.Printf("error in query , reason: %s", err)
	}
	res, err := databuckets.DownloadFromOracle(request.UUID, email)
	if err != nil {
		http.Error(w, "Failed to download file", http.StatusInternalServerError)
		log.Printf("failed to download file , reason: %s", err)

		return
	}

	response := DownloadResponse{Msg: res}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
