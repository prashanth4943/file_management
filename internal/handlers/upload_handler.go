package handlers

import (
	"dms-backend/databuckets"
	"encoding/json"
	"log"
	"net/http"
)

type UploadResponse struct {
	UUID string `json:"uuid"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		FilePath string `json:"file_path"`
		Email    string `json:"email"`
	}

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.FilePath == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload in progress. Please wait..."))

	go func() {
		uuid, err := databuckets.UploadToOracle(input.FilePath, input.Email)
		if err != nil {
			// http.Error(w, "Failed to upload file", http.StatusInternalServerError)
			// log.Printf("failed to upload file , reason: %s", err)
			log.Printf("Failed to upload file for %s, reason: %s", input.Email, err)
			return
		}
		// log.Printf("Upload successful for %s: %s", req.Email, result)
		log.Printf("Upload successful for %s with UUID: %s", input.Email, uuid)

	}()

	// Respond with the URL
	// response := UploadResponse{UUID: uuid}
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}
