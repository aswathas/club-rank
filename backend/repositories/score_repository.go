package repositories

import (
	"backend/models"
	"time"
)

type ScoreRepository interface {
    AddScore(score *models.StudentScore) error
    GetScoresByStudent(studentID uint) ([]models.StudentScore, error)
    GetScoresByDomain(domainID uint) ([]models.StudentScore, error)
    GetScoresByPeriod(start, end time.Time) ([]models.StudentScore, error)
    GetStudentScores(studentID uint) ([]models.StudentScore, error)
}