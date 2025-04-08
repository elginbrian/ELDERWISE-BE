package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type EntityType string

const (
	EntityTypeElder     EntityType = "elder"
	EntityTypeCaregiver EntityType = "caregiver"
	EntityTypeUser      EntityType = "user"
	EntityTypeAgenda    EntityType = "agenda"
	EntityTypeArea      EntityType = "area"
	EntityTypeGeneral   EntityType = "general"
)

func (et EntityType) IsValid() bool {
	switch et {
	case EntityTypeElder, EntityTypeCaregiver, EntityTypeUser, EntityTypeAgenda, EntityTypeArea, EntityTypeGeneral:
		return true
	default:
		return false
	}
}

func (et *EntityType) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch value {
	case "elder", "EntityType.elder":
		*et = EntityTypeElder
	case "caregiver", "EntityType.caregiver":
		*et = EntityTypeCaregiver
	case "user", "EntityType.user":
		*et = EntityTypeUser
	case "agenda", "EntityType.agenda":
		*et = EntityTypeAgenda
	case "area", "EntityType.area":
		*et = EntityTypeArea
	case "general", "EntityType.general", "":
		*et = EntityTypeGeneral
	default:
		return fmt.Errorf("invalid entity type: %s", value)
	}
	return nil
}

type StorageUpload struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Path       string     `json:"path"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UserID     *string    `json:"userId,omitempty"`       // Changed from user_id to match frontend
	EntityID   *string    `json:"entityId,omitempty"`     // Changed from entity_id to match frontend
	EntityType EntityType `json:"entityType"`             // Changed from entity_type to match frontend
}
