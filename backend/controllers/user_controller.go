package controllers

import (
	"net/http"
	"strconv"
	"backend/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{userService: service}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	// ...existing code or stub...
	c.JSON(http.StatusOK, gin.H{"message": "Get all users - stub"})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// ...stub code...
	c.JSON(http.StatusOK, gin.H{"message": "Get user by id", "id": id})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// ...stub code...
	c.JSON(http.StatusCreated, gin.H{"message": "User created - stub"})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	// ...stub code...
	c.JSON(http.StatusOK, gin.H{"message": "User updated - stub"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	// ...stub code...
	c.JSON(http.StatusOK, gin.H{"message": "User deleted - stub"})
}
