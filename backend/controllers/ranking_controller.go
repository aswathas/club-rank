// backend/controllers/ranking_controller.go
package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RankingController struct {
    rankingService services.RankingService
}

func NewRankingController(rankingService services.RankingService) *RankingController {
    return &RankingController{
        rankingService: rankingService,
    }
}

// CalculateRanking godoc
// @Summary Calculate ranking for a club
// @Tags rankings
// @Accept json
// @Produce json
// @Param clubId path int true "Club ID"
// @Success 200 {object} models.Ranking
// @Router /api/rankings/calculate/{clubId} [post]
func (c *RankingController) CalculateRanking(ctx *gin.Context) {
    clubID, err := strconv.ParseUint(ctx.Param("clubId"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid club ID"})
        return
    }

    ranking, err := c.rankingService.CalculateRanking(uint(clubID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, ranking)
}

// GetAllRankings godoc
// @Summary Get all club rankings
// @Tags rankings
// @Produce json
// @Success 200 {array} models.ClubRanking
// @Router /api/rankings [get]
func (c *RankingController) GetAllRankings(ctx *gin.Context) {
    rankings, err := c.rankingService.GetAllRankings()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, rankings)
}

// UpdateCriteria godoc
// @Summary Update ranking criteria
// @Tags rankings
// @Accept json
// @Produce json
// @Param criteria body models.RankingCriteria true "Ranking criteria"
// @Success 200 {object} models.Response
// @Router /api/rankings/criteria [put]
func (c *RankingController) UpdateCriteria(ctx *gin.Context) {
    var criteria models.RankingCriteria
    if err := ctx.ShouldBindJSON(&criteria); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.rankingService.UpdateCriteria(&criteria); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Criteria updated successfully"})
}

// GetCriteria godoc
// @Summary Get ranking criteria
// @Tags rankings
// @Produce json
// @Success 200 {array} models.RankingCriteria
// @Router /api/rankings/criteria [get]
func (c *RankingController) GetCriteria(ctx *gin.Context) {
    criteria, err := c.rankingService.GetCriteria()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, criteria)
}