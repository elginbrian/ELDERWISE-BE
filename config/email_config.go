package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type EmailConfig struct {
	Provider         string 
	FallbackProvider string 

	Host       string
	Port       string
	Username   string
	Password   string
	
	FromEmail  string
	FromName   string
	Secure     bool
	
	SendGridAPIKey string
	
	MailgunAPIKey  string
	MailgunDomain  string
	
	MaxRetries       int
	HealthCheckTimeout time.Duration 
}

func NewEmailConfig() *EmailConfig {
	provider := strings.ToLower(os.Getenv("EMAIL_PROVIDER"))
	if provider == "" || provider == "mock" {
		if os.Getenv("EMAIL_USERNAME") != "" && os.Getenv("EMAIL_PASSWORD") != "" {
			provider = "smtp"
			log.Println("Using SMTP provider based on available credentials")
		} else if os.Getenv("SENDGRID_API_KEY") != "" {
			provider = "sendgrid"
			log.Println("Using SendGrid provider based on available credentials")
		} else if os.Getenv("MAILGUN_API_KEY") != "" && os.Getenv("MAILGUN_DOMAIN") != "" {
			provider = "mailgun"
			log.Println("Using Mailgun provider based on available credentials")
		} else {
			provider = "smtp"
			log.Println("ERROR: No email provider credentials found. SMTP will be used but will likely fail.")
		}
	}
	
	fallback := strings.ToLower(os.Getenv("EMAIL_FALLBACK_PROVIDER"))
	
	host := os.Getenv("EMAIL_HOST")
	if host == "" && provider == "smtp" {
		host = "smtp.gmail.com" 
	}
	
	port := os.Getenv("EMAIL_PORT")
	if port == "" && provider == "smtp" {
		port = "465"
	}
	
	fromEmail := os.Getenv("EMAIL_FROM")
	if fromEmail == "" {
		fromEmail = os.Getenv("EMAIL_USERNAME")
	}
	
	fromName := os.Getenv("EMAIL_FROM_NAME")
	if fromName == "" {
		fromName = "Elderwise Alert System" 
	}
	
	maxRetries := 3
	if val := os.Getenv("EMAIL_MAX_RETRIES"); val != "" {
		fmt.Sscanf(val, "%d", &maxRetries)
	}
	
	config := &EmailConfig{
		Provider:         provider,
		FallbackProvider: fallback,
		
		Host:             host,
		Port:             port,
		Username:         os.Getenv("EMAIL_USERNAME"),
		Password:         os.Getenv("EMAIL_PASSWORD"),
		
		FromEmail:        fromEmail,
		FromName:         fromName,
		Secure:           os.Getenv("EMAIL_SECURE") != "false",
		
		SendGridAPIKey:   os.Getenv("SENDGRID_API_KEY"),
		
		MailgunAPIKey:    os.Getenv("MAILGUN_API_KEY"),
		MailgunDomain:    os.Getenv("MAILGUN_DOMAIN"),
		
		MaxRetries:       maxRetries,
		HealthCheckTimeout: 5 * time.Second, // Default to 5 seconds
	}
	
	log.Printf("Email configuration: Provider=%s, Fallback=%s, From=%s <%s>", 
		config.Provider, 
		config.FallbackProvider,
		config.FromName,
		config.FromEmail)
	
	return config
}

func (c *EmailConfig) ValidateConfig() error {
	switch c.Provider {
	case "smtp":
		if c.Host == "" || c.Port == "" || c.Username == "" || c.Password == "" {
			return fmt.Errorf("incomplete SMTP configuration")
		}
	case "sendgrid":
		if c.SendGridAPIKey == "" {
			return fmt.Errorf("SendGrid API key not provided")
		}
	case "mailgun":
		if c.MailgunAPIKey == "" || c.MailgunDomain == "" {
			return fmt.Errorf("incomplete Mailgun configuration")
		}
	case "mock":
		return fmt.Errorf("mock provider is not allowed")
	default:
		return fmt.Errorf("unknown email provider: %s", c.Provider)
	}
	
	if c.FromEmail == "" {
		return fmt.Errorf("sender email address not provided")
	}
	
	return nil
}
