package repositories

import (
	"backend/models"
	"gorm.io/gorm"
)

// StudentRepository interface defines the methods for student data operations
type StudentRepository interface {
	Create(student *models.Student) error
	GetByID(id uint) (*models.Student, error)
	GetByDomainID(domainID uint) ([]models.Student, error)
	GetStudentsByDomainID(domainID uint) ([]models.Student, error)
	Update(student *models.Student) error
	Delete(id uint) error
}

// StudentRepositoryImpl implements StudentRepository interface
type StudentRepositoryImpl struct {
	db *gorm.DB
}

// NewStudentRepository creates a new instance of StudentRepository
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{db: db}
}

// Create creates a new student record
func (r *StudentRepositoryImpl) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

// GetByID retrieves a student by their ID
func (r *StudentRepositoryImpl) GetByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// GetByDomainID retrieves students by domain ID
func (r *StudentRepositoryImpl) GetByDomainID(domainID uint) ([]models.Student, error) {
	var students []models.Student
	err := r.db.Where("domain_id = ?", domainID).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

// GetStudentsByDomainID retrieves students by domain ID (same as GetByDomainID)
func (r *StudentRepositoryImpl) GetStudentsByDomainID(domainID uint) ([]models.Student, error) {
	return r.GetByDomainID(domainID)
}

// Update updates an existing student record
func (r *StudentRepositoryImpl) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

// Delete removes a student record by ID
func (r *StudentRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}