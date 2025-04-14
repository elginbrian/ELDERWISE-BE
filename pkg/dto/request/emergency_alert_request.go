package dto

import (
	"time"
)

type EmergencyAlertRequestDTO struct {
	ElderID     string    `json:"elder_id" validate:"required"`
	CaregiverID string    `json:"caregiver_id" validate:"required"`
	Datetime    time.Time `json:"datetime" validate:"required"`
	ElderLat    float64   `json:"elder_lat" validate:"required"`
	ElderLong   float64   `json:"elder_long" validate:"required"`
	IsDismissed bool      `json:"is_dismissed default:false"` 
}



