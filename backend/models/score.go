package models

import "time"

type Score struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StudentID uint      `json:"student_id" gorm:"not null"`
	DomainID  uint      `json:"domain_id" gorm:"not null"`
	Score     int       `json:"score" gorm:"not null"`
	Month     time.Time `json:"month" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}