package dto

import (
	"time"
)

type AgendaRequestDTO struct {
	ElderID     string    `json:"elder_id" validate:"required"`
	CaregiverID string    `json:"caregiver_id" validate:"required"`
	Category    string    `json:"category" validate:"required"`
	Content1    string    `json:"content1" validate:"required"`
	Content2    string    `json:"content2,omitempty"`
	Datetime    time.Time `json:"datetime" validate:"required"`
	IsFinished  bool      `json:"is_finished"`
}



