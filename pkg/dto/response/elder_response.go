package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type EldersResponseDTO struct {
	Elders []models.Elder `json:"elders"`
}

type ElderResponseDTO struct {
	Elder models.Elder `json:"elder"`
}
