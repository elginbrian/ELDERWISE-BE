package dto

import "github.com/elginbrian/ELDERWISE-BE/internal/models"

type UserResponseDTO struct {
	User models.User `json:"user"`
}

type UsersResponseDTO struct {
	Users []models.User `json:"users"`
}