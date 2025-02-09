package models

import "time"

// Domain represents a domain in the club ranking system
// @Description Domain model representing areas of expertise or departments within a club
type Domain struct {
    ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
    Name      string    `json:"name" gorm:"not null" example:"Technology" binding:"required"`
    ClubID    uint      `json:"club_id" gorm:"not null" example:"1" binding:"required"`
    CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" example:"2023-01-01T00:00:00Z"`
}

// DomainResponse represents the response format for a single domain
// @Description Response wrapper for a single domain
type DomainResponse struct {
    Status  string  `json:"status" example:"success"`
    Message string  `json:"message" example:"Domain created successfully"`
    Data    *Domain `json:"data,omitempty"`
}

// DomainListResponse represents the response format for multiple domains
// @Description Response wrapper for multiple domains
type DomainListResponse struct {
    Status  string   `json:"status" example:"success"`
    Message string   `json:"message" example:"Domains retrieved successfully"`
    Data    []Domain `json:"data,omitempty"`
}