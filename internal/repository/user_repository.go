package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(userID string) (*models.User, error)
	FindCaregiversByUserID(userID string) ([]models.Caregiver, error)
	FindEldersByUserID(userID string) ([]models.Elder, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindCaregiversByUserID(userID string) ([]models.Caregiver, error) {
	var caregivers []models.Caregiver
	if err := r.db.Where("user_id = ?", userID).Find(&caregivers).Error; err != nil {
		return nil, err
	}
	return caregivers, nil
}

func (r *userRepositoryImpl) FindEldersByUserID(userID string) ([]models.Elder, error) {
	var elders []models.Elder
	if err := r.db.Where("user_id = ?", userID).Find(&elders).Error; err != nil {
		return nil, err
	}
	return elders, nil
}



