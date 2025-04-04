package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type CaregiverRepository interface {
	FindByID(caregiverID string) (*models.Caregiver, error)
	Create(caregiver *models.Caregiver) error
	Update(caregiver *models.Caregiver) error
}

type caregiverRepositoryImpl struct {
	db *gorm.DB
}

func NewCaregiverRepository(db *gorm.DB) CaregiverRepository {
	return &caregiverRepositoryImpl{db}
}

func (r *caregiverRepositoryImpl) FindByID(caregiverID string) (*models.Caregiver, error) {
	var caregiver models.Caregiver
	if err := r.db.First(&caregiver, "caregiver_id = ?", caregiverID).Error; err != nil {
		return nil, err
	}
	return &caregiver, nil
}

func (r *caregiverRepositoryImpl) Create(caregiver *models.Caregiver) error {
	return r.db.Create(caregiver).Error
}

func (r *caregiverRepositoryImpl) Update(caregiver *models.Caregiver) error {
	return r.db.Save(caregiver).Error
}
