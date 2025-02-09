package integration

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestClubDomainIntegration(t *testing.T) {
    // Setup
    domainRepo := repositories.NewDomainRepository(testDB)
    domainService := services.NewDomainService(domainRepo)
    domainController := controllers.NewDomainController(domainService)

    router := gin.New()
    router.GET("/api/clubs/:clubId/domains", domainController.GetDomainsByClubID)

    t.Run("Get Domains By Club ID", func(t *testing.T) {
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("GET", "/api/clubs/1/domains", nil)
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
    })
}