package models

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"password,omitempty" gorm:"not null"`
    Role      string    `json:"role" gorm:"not null"`
    ClubID    uint      `json:"club_id"`
    CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type UserResponse struct {
    Status  string `json:"status" example:"success"`
    Message string `json:"message,omitempty" example:"User created successfully"`
    Data    *User  `json:"data,omitempty"`
}

type UserListResponse struct {
    Status  string `json:"status" example:"success"`
    Message string `json:"message,omitempty" example:"Users retrieved successfully"`
    Data    []User `json:"data,omitempty"`
}