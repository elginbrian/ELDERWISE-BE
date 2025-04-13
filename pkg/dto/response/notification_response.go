package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type NotificationsResponseDTO struct {
	Notifications []models.Notification `json:"notifications"`
}

type UnreadCountResponseDTO struct {
	Count int64 `json:"count"`
}

