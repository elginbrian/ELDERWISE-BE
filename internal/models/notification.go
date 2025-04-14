package models

import "time"

type NotificationType string

const (
	NotificationTypeAreaBreach     NotificationType = "AREA_BREACH"
	NotificationTypeAgendaOverdue  NotificationType = "AGENDA_OVERDUE"
	NotificationTypeAgendaCompleted NotificationType = "AGENDA_COMPLETED"
	NotificationTypeEmergencyAlert NotificationType = "EMERGENCY_ALERT"
)

type Notification struct {
	NotificationID string           `json:"notification_id" gorm:"primaryKey"`
	ElderID        string           `json:"elder_id" gorm:"index"`
	Type           NotificationType `json:"type"`
	Message        string           `json:"message"`
	Datetime       time.Time        `json:"datetime"`
	IsRead         bool             `json:"is_read" gorm:"default:false"`
	RelatedID      string           `json:"related_id"`
	CreatedAt      time.Time        `json:"created_at"`
}




