package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type LocationHistoryResponseDTO struct {
	LocationHistory models.LocationHistory `json:"location_history"`
}

type LocationHistoryPointsResponseDTO struct {
	Points []models.LocationHistoryPoint `json:"points"`
}


