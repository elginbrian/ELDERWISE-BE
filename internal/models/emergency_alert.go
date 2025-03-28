package models

import "time"

type EmergencyAlert struct {
	EmergencyAlertID string    `json:"emergency_alert_id" gorm:"primaryKey"` 
	ElderID          string    `json:"elder_id" gorm:"index"`               
	CaregiverID      string    `json:"caregiver_id" gorm:"index"`            
	Datetime         time.Time `json:"datetime"`
	ElderLat         float64   `json:"elder_lat"`
	ElderLong        float64   `json:"elder_long"`
	IsDismissed      bool      `json:"is_dismissed"`
}