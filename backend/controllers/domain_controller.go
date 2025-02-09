package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DomainController struct {
    domainService services.DomainService
}

func NewDomainController(domainService services.DomainService) *DomainController {
    return &DomainController{
        domainService: domainService,
    }
}

// CreateDomain godoc
// @Summary Create a new domain
// @Description Create a new domain within the system
// @Tags domains
// @Accept json
// @Produce json
// @Param domain body models.Domain true "Domain object"
// @Success 201 {object} models.DomainResponse
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Router /api/domains [post]
func (c *DomainController) CreateDomain(ctx *gin.Context) {
    var domain models.Domain
    if err := ctx.ShouldBindJSON(&domain); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdDomain, err := c.domainService.CreateDomain(&domain)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, models.DomainResponse{
        Status:  "success",
        Message: "Domain created successfully",
        Data:    createdDomain,
    })
}

// GetAllDomains godoc
// @Summary Get all domains
// @Description Retrieve all domains in the system
// @Tags domains
// @Produce json
// @Success 200 {object} models.DomainListResponse
// @Failure 500 {object} models.Response "Internal Server Error"
// @Router /api/domains [get]
func (c *DomainController) GetAllDomains(ctx *gin.Context) {
    domains, err := c.domainService.GetAllDomains()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, models.DomainListResponse{
        Status:  "success",
        Message: "Domains retrieved successfully",
        Data:    domains,
    })
}

// GetDomainByID godoc
// @Summary Get domain by ID
// @Description Retrieve a single domain by its ID
// @Tags domains
// @Produce json
// @Param id path int true "Domain ID"
// @Success 200 {object} models.DomainResponse
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 404 {object} models.Response "Not Found"
// @Router /api/domains/{id} [get]
func (c *DomainController) GetDomainByID(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    domain, err := c.domainService.GetDomainByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Domain not found"})
        return
    }

    ctx.JSON(http.StatusOK, models.DomainResponse{
        Status:  "success",
        Message: "Domain retrieved successfully",
        Data:    domain,
    })
}

// UpdateDomain godoc
// @Summary Update a domain
// @Description Update domain by its ID
// @Tags domains
// @Accept json
// @Produce json
// @Param id path int true "Domain ID"
// @Param domain body models.Domain true "Domain object"
// @Success 200 {object} models.DomainResponse
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Router /api/domains/{id} [put]
func (c *DomainController) UpdateDomain(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var domain models.Domain
    if err := ctx.ShouldBindJSON(&domain); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedDomain, err := c.domainService.UpdateDomain(uint(id), &domain)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, models.DomainResponse{
        Status:  "success",
        Message: "Domain updated successfully",
        Data:    updatedDomain,
    })
}

// DeleteDomain godoc
// @Summary Delete a domain
// @Description Delete a domain by its ID
// @Tags domains
// @Produce json
// @Param id path int true "Domain ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Router /api/domains/{id} [delete]
func (c *DomainController) DeleteDomain(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := c.domainService.DeleteDomain(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Domain deleted successfully",
    })
}

// GetDomainsByClubID godoc
// @Summary Get domains by club ID
// @Description Retrieve all domains for a specific club
// @Tags domains
// @Produce json
// @Param clubId path int true "Club ID"
// @Success 200 {object} models.DomainListResponse
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Router /api/clubs/{clubId}/domains [get]
func (c *DomainController) GetDomainsByClubID(ctx *gin.Context) {
    clubID, err := strconv.ParseUint(ctx.Param("clubId"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid club ID format"})
        return
    }

    domains, err := c.domainService.GetDomainsByClubID(uint(clubID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, models.DomainListResponse{
        Status:  "success",
        Message: "Domains retrieved successfully",
        Data:    domains,
    })
}