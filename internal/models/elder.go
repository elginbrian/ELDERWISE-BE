package models

import "time"

type Elder struct {
	ElderID    string    `json:"elder_id" gorm:"primaryKey"` 
	UserID     string    `json:"user_id" gorm:"index"`       
	Name       string    `json:"name"`
	Birthdate  time.Time `json:"birthdate"`
	Gender     string    `json:"gender"`
	BodyHeight float64   `json:"body_height"`
	BodyWeight float64   `json:"body_weight"`
	PhotoURL   string    `json:"photo_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}


