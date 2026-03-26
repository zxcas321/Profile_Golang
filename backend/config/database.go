package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateDatabaseIfNotExists() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to postgres:", err)
	}
	defer db.Close()

	dbName := os.Getenv("DB_NAME")

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	err = db.QueryRow(query, dbName).Scan(&exists)
	if err != nil {
		log.Fatal("Failed to check database existence:", err)
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbName))
		if err != nil {
			log.Fatal("Failed to create database:", err)
		}
		log.Printf("Database '%s' created successfully!", dbName)
	} else {
		log.Printf("Database '%s' already exists, skipping.", dbName)
	}
}

func DropDatabase() {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }
    defer db.Close()

    dbName := os.Getenv("DB_NAME")
    _, err = db.Exec(fmt.Sprintf(`DROP DATABASE IF EXISTS "%s"`, dbName))
    if err != nil {
        log.Fatal("Failed to drop database:", err)
    }
    log.Printf("Database '%s' dropped!", dbName)
}

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	log.Println("Database connected!")
	return db
}