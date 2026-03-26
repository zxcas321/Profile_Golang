package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"profile_go/config"
	"profile_go/internal/controller"
	"profile_go/internal/repository"
	"profile_go/internal/route"
	"profile_go/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, using system env")
	}

	// Fresh flag — drop & recreate DB
	fresh := flag.Bool("fresh", false, "Drop and recreate database")
	flag.Parse()

	if *fresh {
		config.DropDatabase()
		config.CreateDatabaseIfNotExists()
	} else {
		config.CreateDatabaseIfNotExists()
	}

	// Connect DB
	db := config.ConnectDB()
	defer db.Close()

	// Run migrations
	config.RunMigrations(db)

	// Wire up layers
	userRepo    := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	authCtrl    := controller.NewAuthController(authService)
	userCtrl    := controller.NewUserController(userService)

	// Setup routes
	mux := http.NewServeMux()
	route.SetupRoutes(mux, authCtrl, userCtrl)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}