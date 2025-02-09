// repositories/score_repository_impl.go

package repositories

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type scoreRepository struct {
    db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) ScoreRepository {
    return &scoreRepository{
        db: db,
    }
}

func (r *scoreRepository) AddScore(score *models.StudentScore) error {
    return r.db.Create(score).Error
}

func (r *scoreRepository) GetScoresByStudent(studentID uint) ([]models.StudentScore, error) {
    var scores []models.StudentScore
    err := r.db.Where("student_id = ?", studentID).Find(&scores).Error
    return scores, err
}

func (r *scoreRepository) GetScoresByDomain(domainID uint) ([]models.StudentScore, error) {
    var scores []models.StudentScore
    err := r.db.Where("domain_id = ?", domainID).Find(&scores).Error
    return scores, err
}

func (r *scoreRepository) GetScoresByPeriod(start, end time.Time) ([]models.StudentScore, error) {
    var scores []models.StudentScore
    err := r.db.Where("created_at BETWEEN ? AND ?", start, end).Find(&scores).Error
    return scores, err
}
func (r *scoreRepository) GetScoresByStudentAndDomain(studentID, domainID uint) ([]models.StudentScore, error) {
    var scores []models.StudentScore
    err := r.db.Where("student_id = ? AND domain_id = ?", studentID, domainID).Find(&scores).Error
    return scores, err
}



func (r *scoreRepository) GetStudentScores(studentID uint) ([]models.StudentScore, error) {
    var scores []models.StudentScore
    err := r.db.Where("student_id = ?", studentID).
        Order("month desc").
        Find(&scores).Error
    return scores, err
}