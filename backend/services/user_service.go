package services

import (
	"backend/models"
	"backend/repositories"
)

type UserService interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user models.User) error {
	return s.userRepo.CreateUser(&user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.userRepo.UpdateUser(&user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
