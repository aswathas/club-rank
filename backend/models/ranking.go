package models

type Ranking struct {
	ID        int `json:"id"`
	ClubID    int `json:"club_id"`
	Score     int `json:"score"`
	CreatedAt int `json:"created_at"`
}