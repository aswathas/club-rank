package routes

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Serve Swagger UI BEFORE applying middleware so it remains public.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Apply middleware
	router.Use(middleware.Auth())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Backend is running",
		})
	})

	db := config.DB

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	clubService := services.NewClubService()
	clubController := controllers.NewClubController(clubService)

	api := router.Group("/api")
	{
		api.GET("/users", userController.GetAllUsers)
		api.GET("/users/:id", userController.GetUserByID)
		api.POST("/users", userController.CreateUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)

		api.POST("/clubs", clubController.CreateClub)
		api.GET("/clubs", clubController.GetAllClubs)
		api.GET("/clubs/:id", clubController.GetClubByID)
		api.PUT("/clubs/:id", clubController.UpdateClub)
		api.DELETE("/clubs/:id", clubController.DeleteClub)
	}

	return router
}

// @Summary Get club list
// @Description Get list of all clubs
// @Tags clubs
// @Accept json
// @Produce json
// @Success 200 {array} models.Club
// @Router /clubs [get]
func getClubs(c *gin.Context) {
    // ...existing code...
}
