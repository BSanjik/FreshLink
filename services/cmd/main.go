package main

import (
	"log"
	"services/config"
	"services/internal/router"
	"services/internal/storage"
)

func main() {
	config.LoadEnv()

	db, err := storage.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := storage.RunMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	app := router.New(db)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
