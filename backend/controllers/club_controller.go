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

// CreateClub godoc
// @Summary Create a new club
// @Description Create a new club with the provided information
// @Tags clubs
// @Accept json
// @Produce json
// @Param club body models.Club true "Club object"
// @Success 201 {object} models.ClubResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/clubs [post]
func (c *ClubController) CreateClub(ctx *gin.Context) {
	var club models.Club
	if err := ctx.ShouldBindJSON(&club); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := c.clubService.CreateClub(club); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.ClubResponse{
		Status:  "success",
		Message: "Club created successfully",
		Data:    &club,
	})
}

// GetAllClubs godoc
// @Summary Get all clubs
// @Description Get list of all clubs
// @Tags clubs
// @Accept json
// @Produce json
// @Success 200 {object} models.ClubListResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/clubs [get]
func (c *ClubController) GetAllClubs(ctx *gin.Context) {
	clubs, err := c.clubService.GetAllClubs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ClubListResponse{
		Status: "success",
		Data:   clubs,
	})
}

// GetClubByID godoc
// @Summary Get a club by ID
// @Description Get a club by ID
// @Tags clubs
// @Produce json
// @Param id path int true "Club ID"
// @Success 200 {object} models.Club
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/clubs/{id} [get]
func (c *ClubController) GetClubByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
		return
	}

	club, err := c.clubService.GetClubByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "error",
			Message: "Club not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, club)
}

// UpdateClub godoc
// @Summary Update a club
// @Description Update a club
// @Tags clubs
// @Accept json
// @Produce json
// @Param id path int true "Club ID"
// @Param club body models.Club true "Club"
// @Success 200 {object} models.Club
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/clubs/{id} [put]
func (c *ClubController) UpdateClub(ctx *gin.Context) {
	var club models.Club
	if err := ctx.ShouldBindJSON(&club); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := c.clubService.UpdateClub(club); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, club)
}

// DeleteClub godoc
// @Summary Delete a club
// @Description Delete a club
// @Tags clubs
// @Param id path int true "Club ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/clubs/{id} [delete]
func (c *ClubController) DeleteClub(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
		return
	}

	if err := c.clubService.DeleteClub(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Club deleted successfully",
	})
}