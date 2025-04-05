package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type EmailService interface {
	SendMessage(to, subject, message string) error
}

type emailService struct {
	config *config.EmailConfig
}

func NewEmailService(config *config.EmailConfig) EmailService {
	return &emailService{
		config: config,
	}
}

func (s *emailService) SendMessage(to, subject, message string) error {
	auth := smtp.PlainAuth(
		"",
		s.config.Username,
		s.config.Password,
		s.config.Host,
	)

	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.config.FromName, s.config.FromEmail)
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	
	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(message)
	
	log.Printf("Sending email to %s with subject: %s", to, subject)
	
	// Gmail requires TLS
	smtpServer := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	
	// Set up TLS config
	tlsConfig := &tls.Config{
		ServerName: s.config.Host,
	}
	
	// Connect to SMTP server
	client, err := smtp.Dial(smtpServer)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer client.Close()
	
	// Start TLS
	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %w", err)
	}
	
	// Authenticate
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}
	
	// Set the sender and recipient
	if err = client.Mail(s.config.FromEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}
	
	// Send the email body
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}
	
	_, err = wc.Write([]byte(msg.String()))
	if err != nil {
		return fmt.Errorf("failed to write email body: %w", err)
	}
	
	err = wc.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}
	
	// Send the QUIT command and close the connection
	err = client.Quit()
	if err != nil {
		return fmt.Errorf("failed to quit connection: %w", err)
	}
	
	log.Printf("Successfully sent email to %s", to)
	return nil
}
