package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type ElderRepository interface {
	FindByID(elderID string) (*models.Elder, error)
	Create(elder *models.Elder) error
	Update(elder *models.Elder) error
	FindByUserID(userID string) ([]models.Elder, error)
}

type elderRepositoryImpl struct {
	db *gorm.DB
}

func NewElderRepository(db *gorm.DB) ElderRepository {
	return &elderRepositoryImpl{db: db}
}

func (r *elderRepositoryImpl) FindByID(elderID string) (*models.Elder, error) {
	var elder models.Elder
	if err := r.db.Where("elder_id = ?", elderID).First(&elder).Error; err != nil {
		return nil, err
	}
	return &elder, nil
}

func (r *elderRepositoryImpl) Create(elder *models.Elder) error {
	return r.db.Create(elder).Error
}

func (r *elderRepositoryImpl) Update(elder *models.Elder) error {
	return r.db.Save(elder).Error
}

func (r *elderRepositoryImpl) FindByUserID(userID string) ([]models.Elder, error) {
	var elders []models.Elder
	if err := r.db.Where("user_id = ?", userID).Find(&elders).Error; err != nil {
		return nil, err
	}
	return elders, nil
}
