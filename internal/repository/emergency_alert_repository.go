package repository

import (
	"github.com/elginbrian/ELDERWISE-BE/internal/models"
	"gorm.io/gorm"
)

type EmergencyAlertRepository interface {
	Create(alert *models.EmergencyAlert) error
	FindByID(alertID string) (*models.EmergencyAlert, error)
	Update(alert *models.EmergencyAlert) error
	FindByElderID(elderID string) ([]models.EmergencyAlert, error)
}

type emergencyAlertRepository struct {
	db *gorm.DB
}

func NewEmergencyAlertRepository(db *gorm.DB) EmergencyAlertRepository {
	return &emergencyAlertRepository{db: db}
}

func (r *emergencyAlertRepository) Create(alert *models.EmergencyAlert) error {
	return r.db.Create(alert).Error
}

func (r *emergencyAlertRepository) FindByID(alertID string) (*models.EmergencyAlert, error) {
	var alert models.EmergencyAlert
	if err := r.db.Where("emergency_alert_id = ?", alertID).First(&alert).Error; err != nil {
		return nil, err
	}
	return &alert, nil
}

func (r *emergencyAlertRepository) Update(alert *models.EmergencyAlert) error {
	return r.db.Save(alert).Error
}

func (r *emergencyAlertRepository) FindByElderID(elderID string) ([]models.EmergencyAlert, error) {
	var alerts []models.EmergencyAlert
	if err := r.db.Where("elder_id = ?", elderID).Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}