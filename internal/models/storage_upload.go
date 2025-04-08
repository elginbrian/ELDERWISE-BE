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

// IsValid checks if the EntityType is one of the predefined types
func (et EntityType) IsValid() bool {
	switch et {
	case EntityTypeElder, EntityTypeCaregiver, EntityTypeUser, EntityTypeAgenda, EntityTypeArea, EntityTypeGeneral:
		return true
	default:
		return false
	}
}

// UnmarshalJSON implements custom unmarshaling for EntityType
func (et *EntityType) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Handle the case where frontend enum values might come in different formats
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
	CreatedAt  *time.Time `json:"created_at"`
	UserID     *string    `json:"user_id"`
	EntityID   *string    `json:"entity_id"`
	EntityType EntityType `json:"entity_type"`
}
