package config

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// InitDB initializes the database connection and returns it
func InitDB() (*gorm.DB, error) {
    dsn := "host=172.27.132.255 user=aswath password='1234' dbname=club_rank port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
        return nil, err
    }
    log.Println("Database connected successfully")
    return db, nil
}