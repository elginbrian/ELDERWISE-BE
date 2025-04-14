package services

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
)

type UserService interface {
	GetUserByID(userID string) (*models.User, error)
	GetCaregiversByUserID(userID string) ([]models.Caregiver, error)
	GetEldersByUserID(userID string) ([]models.Elder, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetUserByID(userID string) (*models.User, error) {
	return s.repo.FindByID(userID)
}

func (s *userServiceImpl) GetCaregiversByUserID(userID string) ([]models.Caregiver, error) {
	return s.repo.FindCaregiversByUserID(userID)
}

func (s *userServiceImpl) GetEldersByUserID(userID string) ([]models.Elder, error) {
	return s.repo.FindEldersByUserID(userID)
}


