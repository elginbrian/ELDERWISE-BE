package config

import "os"

type EmailConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	FromEmail  string
	FromName   string
	Secure     bool
}

func NewEmailConfig() *EmailConfig {
	return &EmailConfig{
		Host:      "smtp.gmail.com",
		Port:      "587",
		Username:  os.Getenv("EMAIL_USERNAME"),
		Password:  os.Getenv("EMAIL_PASSWORD"),
		FromEmail: os.Getenv("EMAIL_USERNAME"), // Gmail sends from the authenticated user
		FromName:  os.Getenv("EMAIL_FROM_NAME"),
		Secure:    true,
	}
}
