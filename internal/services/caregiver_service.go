package services

import (
	"errors"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
)

type CaregiverService interface {
	GetCaregiverByID(caregiverID string) (*models.Caregiver, error)
	CreateCaregiver(caregiver *models.Caregiver) error
	UpdateCaregiver(caregiverID string, caregiver *models.Caregiver) error
}

type caregiverServiceImpl struct {
	repo repository.CaregiverRepository
}

func NewCaregiverService(repo repository.CaregiverRepository) CaregiverService {
	return &caregiverServiceImpl{repo}
}

func (s *caregiverServiceImpl) GetCaregiverByID(caregiverID string) (*models.Caregiver, error) {
	return s.repo.FindByID(caregiverID)
}

func (s *caregiverServiceImpl) CreateCaregiver(caregiver *models.Caregiver) error {
	caregiver.CreatedAt = time.Now()
	caregiver.UpdatedAt = time.Now()
	return s.repo.Create(caregiver)
}

func (s *caregiverServiceImpl) UpdateCaregiver(caregiverID string, caregiver *models.Caregiver) error {
	existingCaregiver, err := s.repo.FindByID(caregiverID)
	if err != nil {
		return errors.New("caregiver not found")
	}

	existingCaregiver.Name = caregiver.Name
	existingCaregiver.Birthdate = caregiver.Birthdate
	existingCaregiver.Gender = caregiver.Gender
	existingCaregiver.PhoneNumber = caregiver.PhoneNumber
	existingCaregiver.ProfileURL = caregiver.ProfileURL
	existingCaregiver.Relationship = caregiver.Relationship
	existingCaregiver.UpdatedAt = time.Now()

	return s.repo.Update(existingCaregiver)
}
