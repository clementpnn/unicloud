package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erreur: fichier .env requis")
	}

	connStr := fmt.Sprintf(
		"host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
	)

	log.Printf("Connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	return db
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
