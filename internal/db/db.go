package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

var DBInstance *Database

func ConnectDB(dsn string) error {
	conn, err := sql.Open("mysql", dsn)
	// log.Fatal(conn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Database connected!")
	DBInstance = &Database{Conn: conn}
	return nil
}

func RunMigrations() error {
	if DBInstance == nil || DBInstance.Conn == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	migrationFiles := []string{
		"migrations/users.sql",
		// Add other migration files here
	}
	for _, file := range migrationFiles {
		log.Printf("Running migration: %s", file)
		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		if _, err = DBInstance.Conn.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
		log.Printf("Migration %s applied successfully!", file)
	}
	return nil
}
