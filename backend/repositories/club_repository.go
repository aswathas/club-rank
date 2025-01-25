package repositories

import (
    "backend/models"
    "gorm.io/gorm"
)

type ClubRepository interface {
    CreateClub(club *models.Club) error
    GetAllClubs() ([]models.Club, error)
    GetClubByID(id uint) (*models.Club, error)
    UpdateClub(club *models.Club) error
    DeleteClub(id uint) error
}

type clubRepositoryImpl struct {
    db *gorm.DB
}

func NewClubRepository(db *gorm.DB) ClubRepository {
    return &clubRepositoryImpl{db: db}
}

func (r *clubRepositoryImpl) CreateClub(club *models.Club) error {
    return r.db.Create(club).Error
}

func (r *clubRepositoryImpl) GetAllClubs() ([]models.Club, error) {
    var clubs []models.Club
    err := r.db.Find(&clubs).Error
    return clubs, err
}

func (r *clubRepositoryImpl) GetClubByID(id uint) (*models.Club, error) {
    var club models.Club
    err := r.db.First(&club, id).Error
    return &club, err
}

func (r *clubRepositoryImpl) UpdateClub(club *models.Club) error {
    return r.db.Save(club).Error
}

func (r *clubRepositoryImpl) DeleteClub(id uint) error {
    return r.db.Delete(&models.Club{}, id).Error
}