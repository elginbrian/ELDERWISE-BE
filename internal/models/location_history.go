package models

import "time"

type LocationHistory struct {
	LocationHistoryID string                   `json:"location_history_id" gorm:"primaryKey"` 
	ElderID           string                   `json:"elder_id" gorm:"index"`               
	CaregiverID       string                   `json:"caregiver_id" gorm:"index"`            
	CreatedAt         time.Time                `json:"created_at"`
	Points            []LocationHistoryPoint   `json:"points" gorm:"foreignKey:LocationHistoryID"`
}
