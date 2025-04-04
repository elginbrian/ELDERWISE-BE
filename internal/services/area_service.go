package services

import (
	"time"

	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
)

type AreaService interface {
	GetAreaByID(areaID string) (*models.Area, error)
	CreateArea(area *models.Area) error
	UpdateArea(areaID string, area *models.Area) error
	DeleteArea(areaID string) error
	GetAreasByCaregiver(caregiverID string) ([]models.Area, error)
}

type areaService struct {
	repo repository.AreaRepository
}

func NewAreaService(repo repository.AreaRepository) AreaService {
	return &areaService{repo}
}

func (s *areaService) GetAreaByID(areaID string) (*models.Area, error) {
	return s.repo.FindByID(areaID)
}

func (s *areaService) CreateArea(area *models.Area) error {
	area.CreatedAt = time.Now()
	area.UpdatedAt = time.Now()
	return s.repo.Create(area)
}

func (s *areaService) UpdateArea(areaID string, area *models.Area) error {
	existingArea, err := s.repo.FindByID(areaID)
	if err != nil {
		return err
	}

	area.AreaID = existingArea.AreaID
	area.UpdatedAt = time.Now()
	return s.repo.Update(area)
}

func (s *areaService) DeleteArea(areaID string) error {
	return s.repo.Delete(areaID)
}

func (s *areaService) GetAreasByCaregiver(caregiverID string) ([]models.Area, error) {
	return s.repo.FindByCaregiver(caregiverID)
}
