package routes

import (
	"backend/config"
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Backend is running",
		})
	})

	
	db := config.DB

	
	userRepo := repositories.NewUserRepository(db)

	
	userService := services.NewUserService(userRepo)

	
	userController := controllers.NewUserController(userService)

	
	api := router.Group("/api")
	{
		api.GET("/users", userController.GetAllUsers)
		api.GET("/users/:id", userController.GetUserByID)
		api.POST("/users", userController.CreateUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}

	return router
}
