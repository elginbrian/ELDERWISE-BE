package services

import (
	"fmt"
	"log"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/config"
	"github.com/elginbrian/ELDERWISE-BE/internal/providers"
)

type EmailService interface {
	SendMessage(to, subject, htmlBody string) error
	SendMessageAsync(to, subject, htmlBody string)
	HealthCheck() bool
}

type emailService struct {
	config          *config.EmailConfig
	primaryProvider providers.EmailProvider
	fallbackProvider providers.EmailProvider
}

func createProvider(providerType string, config *config.EmailConfig) (providers.EmailProvider, error) {
	switch providerType {
	case "sendgrid":
		return providers.NewSendGridProvider(config), nil
	case "mailgun":
		return providers.NewMailgunProvider(config), nil
	case "smtp":
		return providers.NewSMTPProvider(config), nil
	default:
		return nil, fmt.Errorf("unknown or unsupported provider type: %s", providerType)
	}
}

func NewEmailService(config *config.EmailConfig) (EmailService, error) {
	if err := config.ValidateConfig(); err != nil {
		return nil, fmt.Errorf("email configuration error: %w", err)
	}
	
	primaryProvider, err := createProvider(config.Provider, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create email provider: %w", err)
	}
	
	var fallbackProvider providers.EmailProvider
	if config.FallbackProvider != "" && config.FallbackProvider != config.Provider {
		fallbackProvider, err = createProvider(config.FallbackProvider, config)
		if err != nil {
			log.Printf("WARNING: Failed to create fallback email provider: %v", err)
		}
	}
	
	return &emailService{
		config:           config,
		primaryProvider:  primaryProvider,
		fallbackProvider: fallbackProvider,
	}, nil
}

// tryWithRetry attempts to send an email with retry mechanism
func (s *emailService) tryWithRetry(provider providers.EmailProvider, to, subject, htmlBody string) error {
	var lastErr error
	
	for attempt := 0; attempt <= s.config.MaxRetries; attempt++ {
		if attempt > 0 {
			// Wait before retry with exponential backoff
			backoff := time.Duration(attempt*attempt) * 500 * time.Millisecond
			time.Sleep(backoff)
			log.Printf("Retrying email to %s (attempt %d/%d)", to, attempt, s.config.MaxRetries)
		}
		
		err := provider.SendEmail(to, subject, htmlBody)
		if err == nil {
			// Success
			if attempt > 0 {
				log.Printf("Email to %s succeeded after %d retries", to, attempt)
			}
			return nil
		}
		
		lastErr = err
		log.Printf("Email attempt failed: %v", err)
	}
	
	return fmt.Errorf("all attempts failed: %w", lastErr)
}

// SendMessage sends an email, falling back if primary provider fails
func (s *emailService) SendMessage(to, subject, htmlBody string) error {
	// Try primary provider with retry
	err := s.tryWithRetry(s.primaryProvider, to, subject, htmlBody)
	if err == nil {
		return nil
	}
	
	// If primary fails and fallback exists, try fallback
	if s.fallbackProvider != nil {
		log.Printf("Primary email provider failed, trying fallback for email to %s", to)
		err = s.tryWithRetry(s.fallbackProvider, to, subject, htmlBody)
		if err == nil {
			return nil
		}
		return fmt.Errorf("both primary and fallback email providers failed: %w", err)
	}
	
	return fmt.Errorf("email delivery failed: %w", err)
}

// SendMessageAsync sends an email asynchronously
func (s *emailService) SendMessageAsync(to, subject, htmlBody string) {
	go func() {
		if err := s.SendMessage(to, subject, htmlBody); err != nil {
			log.Printf("Async email delivery failed: %v", err)
		}
	}()
}

// HealthCheck tests if the email service is functional
func (s *emailService) HealthCheck() bool {
	switch s.config.Provider {
	case "smtp":
		// For SMTP, we can test connection to the server
		provider := s.primaryProvider.(*providers.SMTPProvider)
		return provider.TestConnection() == nil
		
	case "sendgrid", "mailgun":
		// For API-based providers, we consider them healthy if they're configured
		return s.config.ValidateConfig() == nil
		
	default:
		return false
	}
}