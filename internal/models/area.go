package models

import "time"

type Area struct {
	AreaID          string    `json:"area_id" gorm:"primaryKey"` 
	ElderID         string    `json:"elder_id" gorm:"index"`     
	CaregiverID     string    `json:"caregiver_id" gorm:"index"` 
	CenterLat       float64   `json:"center_lat"`
	CenterLong      float64   `json:"center_long"`
	FreeAreaRadius  int       `json:"free_area_radius"`
	WatchAreaRadius int       `json:"watch_area_radius"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

