package db

import (
	"database/sql"
	"fmt"
	"time"

	// "dms-backend/internal/encdec"
	"dms-backend/internal/models"
)

// type Database struct {
// 	Conn *sql.DB
// }

// func (d *Database) CreateUser(user *models.User) error {
func CreateUser(user *models.User) error {
	// log.Fatal(d)
	if DBInstance == nil || DBInstance.Conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	expiryDate := time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02 15:04:05")
	query := `INSERT INTO users (first_name, last_name, email, phone , expiry_date) VALUES (?, ?, ?, ? ,?)`
	result, err := DBInstance.Conn.Exec(query, user.FirstName, user.LastName, user.Email, user.Phone, expiryDate)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	// Get the last inserted ID (optional)
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get insertID: %w", err)
	}
	user.ID = int(id)
	// log.Fatal(user.ID)
	return nil
}

func GetUser(email *models.Input) (string, string, error) {
	if DBInstance == nil || DBInstance.Conn == nil {
		return "", "", fmt.Errorf("database connection is not initialized")
	}
	var firstName, lastName string

	query := "SELECT first_name, last_name FROM users WHERE email = ?"
	err := DBInstance.Conn.QueryRow(query, email.Email).Scan(&firstName, &lastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("no user found with email: %s", email.Email)
		}
		return "", "", fmt.Errorf("failed to get user: %w", err)
	}

	return firstName, lastName, nil
}
