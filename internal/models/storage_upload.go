package models

import "time"

type EntityType string

const (
	EntityTypeElder     EntityType = "elder"
	EntityTypeCaregiver EntityType = "caregiver"
	EntityTypeUser      EntityType = "user"
	EntityTypeAgenda    EntityType = "agenda"
	EntityTypeArea      EntityType = "area"
	EntityTypeGeneral   EntityType = "general"
)

type StorageUpload struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Path       string     `json:"path"`
	CreatedAt  *time.Time `json:"created_at"`
	UserID     *string    `json:"user_id"`
	EntityID   *string    `json:"entity_id"`
	EntityType *string    `json:"entity_type"`
}
