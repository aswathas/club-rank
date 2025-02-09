package models

import "time"

type Student struct {
    ID       uint      `json:"id" gorm:"primaryKey"`
    Name     string    `json:"name" gorm:"not null"`
    DomainID uint      `json:"domain_id" gorm:"not null"`
    JoinedAt time.Time `json:"joined_at" gorm:"not null"`
}

type StudentScore struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    StudentID uint      `json:"student_id" gorm:"not null"`
    Score     float64   `json:"score" gorm:"not null"`
    Month     time.Time `json:"month" gorm:"not null"`
    DomainID  uint      `json:"domain_id" gorm:"not null"`
}