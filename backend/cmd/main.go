package main

import (
	"club-rank/config"
	"club-rank/docs"
	"club-rank/routes"
	"log"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
    // Initialize DB
    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Initialize router
    router := routes.SetupRouter(db)

    // Swagger documentation
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Start server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}