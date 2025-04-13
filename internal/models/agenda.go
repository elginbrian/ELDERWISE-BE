package models

import "time"

type Agenda struct {
	AgendaID    string    `json:"agenda_id" gorm:"primaryKey"` 
	ElderID     string    `json:"elder_id" gorm:"index"`       
	CaregiverID string    `json:"caregiver_id" gorm:"index"`   
	Category    string    `json:"category"`
	Content1    string    `json:"content1"`
	Content2    string    `json:"content2"`
	Datetime    time.Time `json:"datetime"`
	IsFinished  bool      `json:"is_finished"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

