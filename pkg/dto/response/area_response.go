package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type AreaResponseDTO struct {
	Area models.Area `json:"area"`
}

type AreasResponseDTO struct {
	Areas []models.Area `json:"areas"`
}