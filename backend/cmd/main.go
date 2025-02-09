package main

import (
	"backend/config"
	_ "backend/docs" // Ensure your Swagger docs are generated here
	"backend/routes"
	"log"
)

// @title Club Ranking System API
// @version 1.0
// @description API for managing club rankings, domains, and scoring systems
// @host localhost:8080
// @BasePath /
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
