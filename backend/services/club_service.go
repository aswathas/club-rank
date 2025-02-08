package services

import (
	"backend/models"
	"backend/config"
)

type ClubService interface {
	CreateClub(club models.Club) error
	GetAllClubs() ([]models.Club, error)
	GetClubByID(id uint) (models.Club, error)
	UpdateClub(club models.Club) error
	DeleteClub(id uint) error
}

type clubService struct{}

func NewClubService() ClubService {
	return &clubService{}
}

func (s *clubService) CreateClub(club models.Club) error {
	return config.DB.Create(&club).Error
}

func (s *clubService) GetAllClubs() ([]models.Club, error) {
	var clubs []models.Club
	err := config.DB.Find(&clubs).Error
	return clubs, err
}

func (s *clubService) GetClubByID(id uint) (models.Club, error) {
	var club models.Club
	err := config.DB.First(&club, id).Error
	return club, err
}

func (s *clubService) UpdateClub(club models.Club) error {
	return config.DB.Save(&club).Error
}

func (s *clubService) DeleteClub(id uint) error {
	return config.DB.Delete(&models.Club{}, id).Error
}