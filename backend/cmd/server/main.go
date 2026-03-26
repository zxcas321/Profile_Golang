package main

import (
	"log"

	"profile_go/config"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file, using system env")
	}

	// 1. Create DB if not exists
	config.CreateDatabaseIfNotExists()

	// 2. Connect to DB
	db := config.ConnectDB()
	defer db.Close()

	// 3. Run migrations automatically
	config.RunMigrations(db)

	log.Println("Server starting...")
	// next: start HTTP server here
}
