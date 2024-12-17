package routes

import (
	"dms-backend/internal/handlers"
	"net/http"
)

func SetupRoutes() {

	http.HandleFunc("/v1/create_user", handlers.CreateUserHandler)
	http.HandleFunc("/v1/get_user", handlers.GetUserByEmail)
	http.HandleFunc("/v1/upload_document", handlers.UploadHandler)
	http.HandleFunc("/v1/download_document", handlers.DownloadHandler)
	http.HandleFunc("/v1/delete_document", handlers.DeleteHandler)
	// http.HandleFunc("v1/get_ducument", handlers.get_document)
	// http.HandleFunc("v1/delete_document", handlers.delete_document)
	// You can add more routes as needed
}
