package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}
	
	migrationPath := getMigrationPath()

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		os.Getenv("DB_NAME"),
		driver,
	)
	if err != nil {
		log.Fatal("Failed to initialize migrations:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Migrations applied successfully!")
}

func getMigrationPath() string {
    dir, _ := os.Getwd()
    path := filepath.Join(dir, "migration")
    path = filepath.ToSlash(path)
    return "file://" + path
}
