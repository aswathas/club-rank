package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// SetupRouter initializes the routes for the application
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Serve swagger.yaml as a static file to load custom documentation.
	router.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	// Serve Swagger UI (public, no middleware)
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		URL: "/swagger.yaml", // URL for the Swagger documentation
	}, swaggerfiles.Handler))

	// Health check endpoint (public, no middleware)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Backend is running",
		})
	})

	// Initialize repositories, services, and controllers for users
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Initialize services and controllers for clubs
	clubService := services.NewClubService(db)
	clubController := controllers.NewClubController(clubService)

	// Apply middleware to the routes (auth middleware is applied here)
	api := router.Group("/api") // Use "/api" as the base path
	api.Use(middleware.Auth())  // Apply Auth middleware to all routes under /api

	// User routes
	{
		api.GET("/users", userController.GetAllUsers)       // Get all users
		api.GET("/users/:id", userController.GetUserByID)   // Get a single user by ID
		api.POST("/users", userController.CreateUser)       // Create a new user
		api.PUT("/users/:id", userController.UpdateUser)    // Update an existing user by ID
		api.DELETE("/users/:id", userController.DeleteUser) // Delete a user by ID
	}

	// Club routes
	{
		api.POST("/clubs", clubController.CreateClub)       // Create a new club
		api.GET("/clubs", clubController.GetAllClubs)       // Get all clubs
		api.GET("/clubs/:id", clubController.GetClubByID)   // Get a single club by ID
		api.PUT("/clubs/:id", clubController.UpdateClub)    // Update an existing club by ID
		api.DELETE("/clubs/:id", clubController.DeleteClub) // Delete a club by ID
	}

	return router
}
