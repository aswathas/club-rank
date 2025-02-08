package main

import (
	"backend/config"
	"backend/routes"
	"log"
	_ "backend/docs" // Remove space and fix import
)

// @title Club Rank API
// @version 1.0
// @description This is a sample server for Club Rank.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

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
