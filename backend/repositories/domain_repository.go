package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type DomainRepository interface {
    Create(domain *models.Domain) (*models.Domain, error)
    GetAll() ([]models.Domain, error)
    GetByID(id uint) (*models.Domain, error)
    Update(domain *models.Domain) (*models.Domain, error)
    Delete(id uint) error
    GetByClubID(clubID uint) ([]models.Domain, error)
}

type domainRepository struct {
    db *gorm.DB
}

func NewDomainRepository(db *gorm.DB) DomainRepository {
    return &domainRepository{db: db}
}

func (r *domainRepository) Create(domain *models.Domain) (*models.Domain, error) {
    result := r.db.Create(domain)
    return domain, result.Error
}

func (r *domainRepository) GetAll() ([]models.Domain, error) {
    var domains []models.Domain
    result := r.db.Find(&domains)
    return domains, result.Error
}

func (r *domainRepository) GetByID(id uint) (*models.Domain, error) {
    var domain models.Domain
    result := r.db.First(&domain, id)
    return &domain, result.Error
}

func (r *domainRepository) Update(domain *models.Domain) (*models.Domain, error) {
    result := r.db.Save(domain)
    return domain, result.Error
}

func (r *domainRepository) Delete(id uint) error {
    return r.db.Delete(&models.Domain{}, id).Error
}

func (r *domainRepository) GetByClubID(clubID uint) ([]models.Domain, error) {
    var domains []models.Domain
    result := r.db.Where("club_id = ?", clubID).Find(&domains)
    return domains, result.Error
}