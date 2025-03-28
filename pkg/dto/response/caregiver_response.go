package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type CaregiversResponseDTO struct {
	Caregivers []models.Caregiver `json:"caregivers"`
}

type CaregiverResponseDTO struct {
	Caregiver models.Caregiver `json:"caregiver"`
}