package dto

import (
	"time"
)

type LocationHistoryRequestDTO struct {
	ElderID     string    `json:"elder_id" validate:"required"`
	CaregiverID string    `json:"caregiver_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type LocationHistoryPointRequestDTO struct {
	LocationHistoryID string    `json:"location_history_id" validate:"required"`
	Latitude          float64   `json:"latitude" validate:"required"`
	Longitude         float64   `json:"longitude" validate:"required"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}


