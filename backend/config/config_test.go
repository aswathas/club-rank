package config

import (
	"os"
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