// models/ranking.go
package models

import "time"

// Ranking represents a score entry for a club
// @Description Ranking model for storing club scores based on different criteria
type Ranking struct {
    ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
    ClubID    uint      `json:"club_id" gorm:"not null" example:"1"`
    Score     float64   `json:"score" gorm:"not null" example:"85.5"`
    Criteria  string    `json:"criteria" gorm:"not null" example:"Technical Skills"`
    CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
    UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// RankingCriteria represents evaluation criteria for ranking
// @Description Criteria used for evaluating and ranking clubs
type RankingCriteria struct {
    ID          uint    `json:"id" gorm:"primaryKey" example:"1"`
    Name        string  `json:"name" gorm:"not null" example:"Technical Skills"`
    Weight      float64 `json:"weight" gorm:"not null" example:"0.4"`
    Description string  `json:"description" example:"Evaluation of technical capabilities"`
}

// ClubRanking represents the overall ranking of a club
// @Description Overall ranking information for a club
type ClubRanking struct {
    ClubID       uint    `json:"club_id" example:"1"`
    ClubName     string  `json:"club_name" example:"Tech Club"`
    OverallScore float64 `json:"overall_score" example:"87.5"`
    Rank         int     `json:"rank" example:"1"`
}