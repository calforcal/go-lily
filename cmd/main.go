package main

import (
	"log"

	"github.com/calforcal/can-lily-eat-it/api"
	"github.com/calforcal/can-lily-eat-it/storage"
	"github.com/calforcal/can-lily-eat-it/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()
	// Initialize database
	store, err := storage.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer storage.CloseDB(store)

	// Initialize Echo
	e := echo.New()

	// Setup routes
	router := api.NewApiRouter(e, store)
	router.RegisterRoutes()

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
