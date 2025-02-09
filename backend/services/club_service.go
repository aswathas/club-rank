package services

import (
	"backend/models"
	"gorm.io/gorm"
)

type ClubService interface {
	CreateClub(club models.Club) error
	GetAllClubs() ([]models.Club, error)
	GetClubByID(id uint) (models.Club, error)
	UpdateClub(club models.Club) error
	DeleteClub(id uint) error
}

type clubService struct {
	db *gorm.DB
}

func NewClubService(db *gorm.DB) ClubService {
	return &clubService{db: db}
}

func (s *clubService) CreateClub(club models.Club) error {
	return s.db.Create(&club).Error
}

func (s *clubService) GetAllClubs() ([]models.Club, error) {
	var clubs []models.Club
	err := s.db.Find(&clubs).Error
	return clubs, err
}

func (s *clubService) GetClubByID(id uint) (models.Club, error) {
	var club models.Club
	err := s.db.First(&club, id).Error
	if err == gorm.ErrRecordNotFound {
		return models.Club{}, err
	}
	return club, err
}

func (s *clubService) UpdateClub(club models.Club) error {
	return s.db.Save(&club).Error
}

func (s *clubService) DeleteClub(id uint) error {
	return s.db.Delete(&models.Club{}, id).Error
}