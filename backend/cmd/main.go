package main

import (
	"backend/routes" // Ensure this matches your project structure
	"log"

	"github.com/labstack/echo/v4" // Import Echo
)

func main() {
	// Initialize Echo router
	e := echo.New()

	// Set up application routes
	routes.SetupRoutes(e)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	e.Start(":8080")
}
