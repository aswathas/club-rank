package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	// Initialize the database connection
	config.InitDB()

	// Initialize the router
	router := routes.SetupRouter()

	// Start the server
	log.Println("Starting server on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
