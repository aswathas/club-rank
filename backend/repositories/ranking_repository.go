package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type RankingRepository interface {
    Create(ranking *models.Ranking) error
    GetAll() ([]models.Ranking, error)
    GetByClubID(clubID uint) (*models.Ranking, error)
    Update(ranking *models.Ranking) error
    Delete(id uint) error
    UpdateCriteria(criteria *models.RankingCriteria) error
    GetCriteria() ([]models.RankingCriteria, error)
}

type rankingRepository struct {
    db *gorm.DB
}

func NewRankingRepository(db *gorm.DB) RankingRepository {
    return &rankingRepository{db: db}
}

func (r *rankingRepository) Create(ranking *models.Ranking) error {
    return r.db.Create(ranking).Error
}

func (r *rankingRepository) GetAll() ([]models.Ranking, error) {
    var rankings []models.Ranking
    err := r.db.Order("score desc").Find(&rankings).Error
    return rankings, err
}

func (r *rankingRepository) GetByClubID(clubID uint) (*models.Ranking, error) {
    var ranking models.Ranking
    err := r.db.Where("club_id = ?", clubID).First(&ranking).Error
    return &ranking, err
}

func (r *rankingRepository) Update(ranking *models.Ranking) error {
    return r.db.Save(ranking).Error
}

func (r *rankingRepository) Delete(id uint) error {
    return r.db.Delete(&models.Ranking{}, id).Error
}

func (r *rankingRepository) UpdateCriteria(criteria *models.RankingCriteria) error {
    return r.db.Save(criteria).Error
}

func (r *rankingRepository) GetCriteria() ([]models.RankingCriteria, error) {
    var criteria []models.RankingCriteria
    err := r.db.Find(&criteria).Error
    return criteria, err
}