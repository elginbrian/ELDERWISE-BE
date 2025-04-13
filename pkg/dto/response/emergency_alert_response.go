package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type EmergencyAlertResponseDTO struct {
	EmergencyAlert models.EmergencyAlert `json:"emergency_alert"`
}

