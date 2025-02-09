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
	domainRepo := repositories.NewDomainRepository(db)
    domainService := services.NewDomainService(domainRepo)
    domainController := controllers.NewDomainController(domainService)

	// Initialize repositories, services, and controllers for users

	// Initialize services and controllers for clubs
	clubService := services.NewClubService(db)
	clubController := controllers.NewClubController(clubService)

	// Initialize repositories and services for rankings
	rankingRepo := repositories.NewRankingRepository(db)
	clubRepo := repositories.NewClubRepository(db)
	domainRepo = repositories.NewDomainRepository(db)
	studentRepo := repositories.NewStudentRepository(db)
	rankingService := services.NewRankingService(rankingRepo, clubRepo, domainRepo, studentRepo)
	rankingController := controllers.NewRankingController(rankingService)

	// Apply middleware to the routes (auth middleware is applied here)
	api := router.Group("/api") // Use "/api" as the base path
	api.Use(middleware.Auth())  // Apply Auth middleware to all routes under /api

	// User routes
	clubs := api.Group("/clubs")
    {
        clubs.POST("", clubController.CreateClub)        // Create a new club
        clubs.GET("", clubController.GetAllClubs)        // Get all clubs
        clubs.GET("/:id", clubController.GetClubByID)    // Get a single club by ID
        clubs.PUT("/:id", clubController.UpdateClub)     // Update an existing club by ID
        clubs.DELETE("/:id", clubController.DeleteClub)  // Delete a club by ID
        
        // Nested domains route under clubs
    }

    // Domain routes
    domains := api.Group("/domains")
    {
        domains.POST("", domainController.CreateDomain)       // Create domain
        domains.GET("", domainController.GetAllDomains)      // Get all domains
        domains.GET("/:id", domainController.GetDomainByID)  // Get single domain
        domains.PUT("/:id", domainController.UpdateDomain)   // Update domain
        domains.DELETE("/:id", domainController.DeleteDomain) // Delete domain
    }

    // Ranking routes
    rankings := api.Group("/rankings")
    {
        rankings.POST("/calculate/:clubId", rankingController.CalculateRanking)
        rankings.GET("", rankingController.GetAllRankings)
        rankings.PUT("/criteria", rankingController.UpdateCriteria)
        rankings.GET("/criteria", rankingController.GetCriteria)
    }

    return router
}
