package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type LocationHistoryRepository struct {
	DB *gorm.DB
}

func NewLocationHistoryRepository(db *gorm.DB) *LocationHistoryRepository {
	return &LocationHistoryRepository{DB: db}
}

func (r *LocationHistoryRepository) GetLocationHistoryByID(historyID string) (models.LocationHistory, error) {
	var history models.LocationHistory
	err := r.DB.Where("location_history_id = ?", historyID).Preload("Points").First(&history).Error
	return history, err
}

func (r *LocationHistoryRepository) GetLocationHistoryPoints(historyID string) ([]models.LocationHistoryPoint, error) {
	var points []models.LocationHistoryPoint
	err := r.DB.Where("location_history_id = ?", historyID).Find(&points).Error
	return points, err
}

func (r *LocationHistoryRepository) GetElderLocationHistory(elderID string) (models.LocationHistory, error) {
	var history models.LocationHistory
	err := r.DB.Where("elder_id = ?", elderID).Preload("Points").Order("created_at DESC").First(&history).Error
	return history, err
}

func (r *LocationHistoryRepository) CreateLocationHistory(history models.LocationHistory) (models.LocationHistory, error) {
	err := r.DB.Create(&history).Error
	return history, err
}

func (r *LocationHistoryRepository) AddLocationPoint(point models.LocationHistoryPoint) error {
	return r.DB.Create(&point).Error
}