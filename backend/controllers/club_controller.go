package controllers

import (
    "net/http"
    "strconv"
    "backend/models"
    "backend/services"
    "github.com/gin-gonic/gin"
)

type ClubController struct {
    clubService services.ClubService
}

func NewClubController(clubService services.ClubService) *ClubController {
    return &ClubController{clubService: clubService}
}

func (c *ClubController) CreateClub(ctx *gin.Context) {
    var club models.Club
    if err := ctx.ShouldBindJSON(&club); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.clubService.CreateClub(club); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, club)
}

func (c *ClubController) GetAllClubs(ctx *gin.Context) {
    clubs, err := c.clubService.GetAllClubs()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, clubs)
}

func (c *ClubController) GetClubByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    club, err := c.clubService.GetClubByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Club not found"})
        return
    }

    ctx.JSON(http.StatusOK, club)
}

func (c *ClubController) UpdateClub(ctx *gin.Context) {
    var club models.Club
    if err := ctx.ShouldBindJSON(&club); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.clubService.UpdateClub(club); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, club)
}

func (c *ClubController) DeleteClub(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := c.clubService.DeleteClub(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Club deleted successfully"})
}