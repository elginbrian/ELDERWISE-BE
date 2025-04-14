package dto

import "time"

type RegisterResponseDTO struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}


