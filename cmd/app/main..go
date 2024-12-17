package main

import (
	"log"
	"net/http"
	"os"

	"dms-backend/internal/db"
	"dms-backend/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Read the database DSN from environment variables

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding with system env variables.")
	}

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		log.Fatal("DB_DSN environment variable is required")
	}

	if err := db.ConnectDB(dsn); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := db.RunMigrations(); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	routes.SetupRoutes()

	log.Println("Application is ready!")

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
