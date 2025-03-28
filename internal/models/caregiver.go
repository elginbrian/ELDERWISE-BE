package models

import "time"

type Caregiver struct {
	CaregiverID  string    `json:"caregiver_id" gorm:"primaryKey"` 
	UserID       string    `json:"user_id" gorm:"index"`           
	Name         string    `json:"name"`
	Birthdate    time.Time `json:"birthdate"`
	Gender       string    `json:"gender"`
	PhoneNumber  string    `json:"phone_number"`
	ProfileURL   string    `json:"profile_url"`
	Relationship string    `json:"relationship"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}