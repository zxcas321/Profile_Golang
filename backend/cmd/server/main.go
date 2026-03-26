package main

import (
	"log"
	"flag"

	"profile_go/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, using system env")
	}

    fresh := flag.Bool("fresh", false, "Drop and recreate database")
    flag.Parse()

    if *fresh {
        config.DropDatabase()
        config.CreateDatabaseIfNotExists()
    }

	db := config.ConnectDB()
	defer db.Close()

	config.RunMigrations(db)

	log.Println("Server starting...")
}
