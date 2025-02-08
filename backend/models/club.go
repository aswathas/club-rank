package models

import "time"

// Club represents a club entity
type Club struct {
	ID          uint      `json:"id" example:"1"`
	CreatedAt   time.Time `json:"created_at" example:"2025-02-08T16:13:44Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-02-08T16:13:44Z"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" swaggertype:"string" example:"null"`
	Name        string    `json:"name" example:"Chess Club"`
	Description string    `json:"description" example:"A club for chess enthusiasts"`
	Category    string    `json:"category" example:"Sports"`
	Rating      float64   `json:"rating" example:"4.5"`
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