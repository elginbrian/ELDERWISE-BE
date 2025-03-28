package models

import "time"

type LocationHistoryPoint struct {
	PointID           string    `json:"point_id" gorm:"primaryKey"` 
	LocationHistoryID string    `json:"location_history_id" gorm:"index"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	Timestamp         time.Time `json:"timestamp"`
}