package services

import (
    "backend/models"
    "backend/repositories"
)

type ClubService interface {
    CreateClub(club models.Club) error
    GetAllClubs() ([]models.Club, error)
    GetClubByID(id uint) (*models.Club, error)
    UpdateClub(club models.Club) error
    DeleteClub(id uint) error
}

type clubService struct {
    clubRepo repositories.ClubRepository
}

func NewClubService(clubRepo repositories.ClubRepository) ClubService {
    return &clubService{clubRepo: clubRepo}
}

func (s *clubService) CreateClub(club models.Club) error {
    return s.clubRepo.CreateClub(&club)
}

func (s *clubService) GetAllClubs() ([]models.Club, error) {
    return s.clubRepo.GetAllClubs()
}

func (s *clubService) GetClubByID(id uint) (*models.Club, error) {
    return s.clubRepo.GetClubByID(id)
}

func (s *clubService) UpdateClub(club models.Club) error {
    return s.clubRepo.UpdateClub(&club)
}

func (s *clubService) DeleteClub(id uint) error {
    return s.clubRepo.DeleteClub(id)
}