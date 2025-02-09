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

// SetupRouter initializes the routes for the application.
func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    // Serve swagger.yaml and Swagger UI publicly.
    router.StaticFile("/swagger.yaml", "./docs/swagger.yaml")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    // Health check endpoint.
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "Backend is running"})
    })

    // Initialize repositories and services.
    domainRepo := repositories.NewDomainRepository(db)
    domainService := services.NewDomainService(domainRepo)

    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    // Pass DomainService to DomainController.
    domainController := controllers.NewDomainController(domainService)

    clubService := services.NewClubService(db)
    clubController := controllers.NewClubController(clubService)
    clubRepo := repositories.NewClubRepository(db)

    rankingRepo := repositories.NewRankingRepository(db)
    // Reinitialize domain repo for ranking if needed.
    domainRepoForRanking := repositories.NewDomainRepository(db)
    studentRepo := repositories.NewStudentRepository(db)
    rankingService := services.NewRankingService(rankingRepo, clubRepo, domainRepoForRanking, studentRepo)
    rankingController := controllers.NewRankingController(rankingService)

    // API routes group with authentication middleware.
    api := router.Group("/api")
    api.Use(middleware.Auth())

    // User routes.
    users := api.Group("/users")
    {
        users.GET("", userController.GetAllUsers)
        users.GET("/:id", userController.GetUserByID)
        users.POST("", userController.CreateUser)
        users.PUT("/:id", userController.UpdateUser)
        users.DELETE("/:id", userController.DeleteUser)
    }

    // Club routes.
    clubs := api.Group("/clubs")
    {
        clubs.POST("", clubController.CreateClub)
        clubs.GET("", clubController.GetAllClubs)
        clubs.PUT("/:id", clubController.UpdateClub)
        clubs.DELETE("/:id", clubController.DeleteClub)

        // Group club detail routes.
        clubDetail := clubs.Group("/:id")
        {
            clubDetail.GET("", clubController.GetClubByID)
            clubDetail.GET("/domains", domainController.GetDomainsByClubID)
        }
    }

    // Domain routes.
    domains := api.Group("/domains")
    {
        domains.POST("", domainController.CreateDomain)
        domains.GET("", domainController.GetAllDomains)
        domains.GET("/:id", domainController.GetDomainByID)
        domains.PUT("/:id", domainController.UpdateDomain)
        domains.DELETE("/:id", domainController.DeleteDomain)
    }

    // Ranking routes.
    rankings := api.Group("/rankings")
    {
        rankings.POST("/calculate/:clubId", rankingController.CalculateRanking)
        rankings.GET("", rankingController.GetAllRankings)
        rankings.PUT("/criteria", rankingController.UpdateCriteria)
        rankings.GET("/criteria", rankingController.GetCriteria)
    }

    return router
}