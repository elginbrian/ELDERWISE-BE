package models

import "time"

type User struct {
	UserID    string    `json:"user_id" gorm:"primaryKey"` 
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}



