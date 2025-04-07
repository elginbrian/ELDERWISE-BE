package services

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"github.com/elginbrian/ELDERWISE-BE/internal/repository"
)

type LocationHistoryService struct {
	repo *repository.LocationHistoryRepository
}

func NewLocationHistoryService(repo *repository.LocationHistoryRepository) *LocationHistoryService {
	return &LocationHistoryService{repo: repo}
}

func (s *LocationHistoryService) GetLocationHistoryByID(historyID string) (models.LocationHistory, error) {
	return s.repo.GetLocationHistoryByID(historyID)
}

func (s *LocationHistoryService) GetLocationHistoryPoints(historyID string) ([]models.LocationHistoryPoint, error) {
	return s.repo.GetLocationHistoryPoints(historyID)
}

func (s *LocationHistoryService) GetElderLocationHistory(elderID string) (models.LocationHistory, error) {
	return s.repo.GetElderLocationHistory(elderID)
}

func (s *LocationHistoryService) CreateLocationHistory(history models.LocationHistory) (models.LocationHistory, error) {
	return s.repo.CreateLocationHistory(history)
}

func (s *LocationHistoryService) AddLocationPoint(point models.LocationHistoryPoint) error {
	return s.repo.AddLocationPoint(point)
}