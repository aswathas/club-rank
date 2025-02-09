package integration

import (
	"backend/config"
	"backend/models"
	"log"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
    // Set test environment
    os.Setenv("GO_ENV", "test")
    
    // Initialize test database
    var err error
    testDB, err = config.InitDB()
    if err != nil {
        log.Fatal("Failed to connect to test database:", err)
    }

    // Run migrations
    testDB.AutoMigrate(&models.Domain{}, &models.Club{})

    // Run tests
    code := m.Run()

    // Cleanup
    sqlDB, err := testDB.DB()
    if err == nil {
        sqlDB.Close()
    }

    os.Exit(code)
}