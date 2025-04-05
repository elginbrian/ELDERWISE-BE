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
	port := os.Getenv("EMAIL_PORT")
	if port == "" {
		port = "465"
	}
	
	return &EmailConfig{
		Host:      "smtp.gmail.com",
		Port:      port,
		Username:  os.Getenv("EMAIL_USERNAME"),
		Password:  os.Getenv("EMAIL_PASSWORD"),
		FromEmail: os.Getenv("EMAIL_USERNAME"),
		FromName:  os.Getenv("EMAIL_FROM_NAME"),
		Secure:    true,
	}
}
