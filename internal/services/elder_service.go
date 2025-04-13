package services

import (
	"errors"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
)

type ElderService interface {
	GetElderByID(elderID string) (*models.Elder, error)
	CreateElder(elder *models.Elder) (*models.Elder, error)
	UpdateElder(elderID string, elder *models.Elder) (*models.Elder, error)
	GetEldersByUserID(userID string) ([]models.Elder, error)
}

type elderServiceImpl struct {
	repo repository.ElderRepository
}

func NewElderService(repo repository.ElderRepository) ElderService {
	return &elderServiceImpl{repo: repo}
}

func (s *elderServiceImpl) GetElderByID(elderID string) (*models.Elder, error) {
	return s.repo.FindByID(elderID)
}

func (s *elderServiceImpl) CreateElder(elder *models.Elder) (*models.Elder, error) {
	elder.CreatedAt = time.Now()
	elder.UpdatedAt = time.Now()
	if err := s.repo.Create(elder); err != nil {
		return nil, err
	}
	return elder, nil
}

func (s *elderServiceImpl) UpdateElder(elderID string, elder *models.Elder) (*models.Elder, error) {
	existingElder, err := s.repo.FindByID(elderID)
	if err != nil {
		return nil, errors.New("elder not found")
	}

	existingElder.Name = elder.Name
	existingElder.Birthdate = elder.Birthdate
	existingElder.Gender = elder.Gender
	existingElder.BodyHeight = elder.BodyHeight
	existingElder.BodyWeight = elder.BodyWeight
	existingElder.PhotoURL = elder.PhotoURL
	existingElder.UpdatedAt = time.Now()

	if err := s.repo.Update(existingElder); err != nil {
		return nil, err
	}
	return existingElder, nil
}

func (s *elderServiceImpl) GetEldersByUserID(userID string) ([]models.Elder, error) {
	return s.repo.FindByUserID(userID)
}


