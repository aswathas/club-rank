package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestConfig struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
}

func LoadTestConfig() *TestConfig {
    return &TestConfig{
        DBHost:     getEnvOrDefault("TEST_DB_HOST", "172.27.132.255"),
        DBUser:     getEnvOrDefault("TEST_DB_USER", "aswath"),
        DBPassword: getEnvOrDefault("TEST_DB_PASSWORD", "1234"),
        DBName:     getEnvOrDefault("TEST_DB_NAME", "club_rank"),
        DBPort:     getEnvOrDefault("TEST_DB_PORT", "5432"),
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func SetupTestDB() (*gorm.DB, error) {
    config := LoadTestConfig()
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        config.DBHost,
        config.DBUser,
        config.DBPassword,
        config.DBName,
        config.DBPort,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to test database: %v", err)
    }

    return db, nil
}

func CleanupTestDB(db *gorm.DB) error {
    sqlDB, err := db.DB()
    if err != nil {
        return err
    }
    return sqlDB.Close()
}