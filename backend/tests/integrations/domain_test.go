package integration

import (
	"backend/controllers"
	"backend/models"
	"backend/repositories"
	"backend/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDomainIntegration(t *testing.T) {
    // Setup
    repo := repositories.NewDomainRepository(testDB)
    service := services.NewDomainService(repo)
    controller := controllers.NewDomainController(service)
    
    router := gin.New()
    router.POST("/api/domains", controller.CreateDomain)
    router.GET("/api/domains", controller.GetAllDomains)

    t.Run("Create and Get Domain", func(t *testing.T) {
        // Test Create Domain
        domain := models.Domain{
            Name:   "Test Domain",
            ClubID: 1,
        }
        
        w := httptest.NewRecorder()
        body, _ := json.Marshal(domain)
        req, _ := http.NewRequest("POST", "/api/domains", bytes.NewBuffer(body))
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
        
        // Test Get All Domains
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("GET", "/api/domains", nil)
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
    })
}