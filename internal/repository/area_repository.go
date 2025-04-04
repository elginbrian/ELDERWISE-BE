package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type AreaRepository interface {
	FindByID(areaID string) (*models.Area, error)
	Create(area *models.Area) error
	Update(area *models.Area) error
	Delete(areaID string) error
	FindByCaregiver(caregiverID string) ([]models.Area, error)
}

type areaRepository struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) AreaRepository {
	return &areaRepository{db}
}

func (r *areaRepository) FindByID(areaID string) (*models.Area, error) {
	var area models.Area
	if err := r.db.First(&area, "area_id = ?", areaID).Error; err != nil {
		return nil, err
	}
	return &area, nil
}

func (r *areaRepository) Create(area *models.Area) error {
	return r.db.Create(area).Error
}

func (r *areaRepository) Update(area *models.Area) error {
	return r.db.Save(area).Error
}

func (r *areaRepository) Delete(areaID string) error {
	return r.db.Delete(&models.Area{}, "area_id = ?", areaID).Error
}

func (r *areaRepository) FindByCaregiver(caregiverID string) ([]models.Area, error) {
	var areas []models.Area
	err := r.db.Where("caregiver_id = ?", caregiverID).Find(&areas).Error
	return areas, err
}
