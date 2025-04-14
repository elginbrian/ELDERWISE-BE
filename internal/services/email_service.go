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

func NewLoggingEmailService() EmailService {
	return &loggingEmailService{}
}

type loggingEmailService struct{}

func (s *loggingEmailService) SendMessage(to, subject, htmlBody string) error {
	log.Printf("LOGGING ONLY (Email not sent): To=%s, Subject=%s, Length=%d", 
		to, subject, len(htmlBody))
	return nil
}

func (s *loggingEmailService) SendMessageAsync(to, subject, htmlBody string) {
	go func() {
		_ = s.SendMessage(to, subject, htmlBody)
	}()
}

func (s *loggingEmailService) HealthCheck() bool {
	return true
}

func (s *emailService) tryWithRetry(provider providers.EmailProvider, to, subject, htmlBody string) error {
	var lastErr error
	
	for attempt := 0; attempt <= s.config.MaxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(attempt*attempt) * 500 * time.Millisecond
			time.Sleep(backoff)
			log.Printf("Retrying email to %s (attempt %d/%d)", to, attempt, s.config.MaxRetries)
		}
		
		err := provider.SendEmail(to, subject, htmlBody)
		if err == nil {
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

func (s *emailService) SendMessage(to, subject, htmlBody string) error {
	err := s.tryWithRetry(s.primaryProvider, to, subject, htmlBody)
	if err == nil {
		return nil
	}
	
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

func (s *emailService) SendMessageAsync(to, subject, htmlBody string) {
	go func() {
		if err := s.SendMessage(to, subject, htmlBody); err != nil {
			log.Printf("Async email delivery failed: %v", err)
		}
	}()
}

func (s *emailService) HealthCheck() bool {
	switch s.config.Provider {
	case "smtp":
		provider := s.primaryProvider.(*providers.SMTPProvider)
		return provider.TestConnection() == nil
		
	case "sendgrid", "mailgun":
		return s.config.ValidateConfig() == nil
		
	default:
		return false
	}
}


