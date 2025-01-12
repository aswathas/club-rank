package routes

import (
    "backend/controllers"
    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    // Example route
    e.GET("/health", controllers.HealthCheck)
}
