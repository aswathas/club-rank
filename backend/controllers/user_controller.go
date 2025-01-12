package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{
        "status": "Backend is running!",
    })
}
