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

	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
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
