package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// Connect menginisialisasi dan mengembalikan koneksi ke database
func Connect() (*sql.DB, error) {
	// Pastikan Anda sudah mengatur environment variables ini di file .env
	// Contoh: DATABASE_URL="postgres://user:password@localhost:5432/komunal_db?sslmode=disable"
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open sql connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}