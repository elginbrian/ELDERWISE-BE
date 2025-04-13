package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type AgendaResponseDTO struct {
	Agenda models.Agenda `json:"agenda"`
}

type AgendasResponseDTO struct {
	Agendas []models.Agenda `json:"agendas"`
}
