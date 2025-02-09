package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	UserID uint `json:"user_id"`
	ClubID uint `json:"club_id"`
	Score  int  `json:"score"`
}