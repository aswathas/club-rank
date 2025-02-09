package services

import (
	"backend/models"
	"backend/repositories"
)

type DomainService interface {
    CreateDomain(domain *models.Domain) (*models.Domain, error)
    GetAllDomains() ([]models.Domain, error)
    GetDomainByID(id uint) (*models.Domain, error)
    UpdateDomain(id uint, domain *models.Domain) (*models.Domain, error)
    DeleteDomain(id uint) error
    GetDomainsByClubID(clubID uint) ([]models.Domain, error)
}

type domainService struct {
    repo repositories.DomainRepository
}

func NewDomainService(repo repositories.DomainRepository) DomainService {
    return &domainService{
        repo: repo,
    }
}

func (s *domainService) CreateDomain(domain *models.Domain) (*models.Domain, error){
	return s.repo.Create(domain)
}
func (s *domainService) GetAllDomains() ([]models.Domain, error) {
    return s.repo.GetAll()
}
func (s *domainService) GetDomainByID(id uint) (*models.Domain, error) {
    return s.repo.GetByID(id)
}
func (s *domainService) UpdateDomain(id uint, domain *models.Domain) (*models.Domain, error) {
    domain.ID = id
    return s.repo.Update(domain)
}	
func (s *domainService) DeleteDomain(id uint) error {
	return s.repo.Delete(id)
}
func (s *domainService) GetDomainsByClubID(clubID uint) ([]models.Domain, error) {
    return s.repo.GetByClubID(clubID)
}
