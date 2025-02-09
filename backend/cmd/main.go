package main

import (
	"backend/config"
	_ "backend/docs" // Ensure your Swagger docs are generated here
	"backend/routes"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a user API.
// @host localhost:8080
// @BasePath /
// main.go
func main() {
	// Initialize DB
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize router and pass the db connection
	router := routes.SetupRouter(db)

	// Start server
	log.Println("Starting server on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
