package models

import (
	"time"
)

type Club struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Logo      string
    Subdomain string
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// ClubResponse represents the response for a club operation
type ClubResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message,omitempty" example:"Club created successfully"`
	Data    *Club  `json:"data,omitempty"`
}

// ClubListResponse represents the response for multiple clubs
type ClubListResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message,omitempty" example:"Clubs retrieved successfully"`
	Data    []Club `json:"data,omitempty"`
}
