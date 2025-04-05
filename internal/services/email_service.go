package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type EmailService interface {
	SendMessage(to, subject, message string) error
	SendMessageAsync(to, subject, message string)
}

type emailService struct {
	config *config.EmailConfig
}

func NewEmailService(config *config.EmailConfig) EmailService {
	return &emailService{
		config: config,
	}
}

func (s *emailService) SendMessageAsync(to, subject, message string) {
	go func() {
		if err := s.SendMessage(to, subject, message); err != nil {
			log.Printf("Error sending email asynchronously: %v", err)
		}
	}()
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
	
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	
	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	
	client, err := smtp.NewClient(conn, s.config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()
	
	if tc, ok := client.TLSConnectionState(); !ok {
		tlsConfig := &tls.Config{
			ServerName: s.config.Host,
			MinVersion: tls.VersionTLS12,
		}
		
		if err = client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("failed to start TLS: %w", err)
		}
	} else {
		log.Printf("Already using TLS version %x", tc.Version)
	}
	
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}
	
	if err = client.Mail(s.config.FromEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}
	
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}
	
	_, err = wc.Write([]byte(msg.String()))
	if err != nil {
		wc.Close()
		return fmt.Errorf("failed to write email body: %w", err)
	}
	
	if err = wc.Close(); err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}
	
	if err = client.Quit(); err != nil {
		return fmt.Errorf("failed to quit connection: %w", err)
	}
	
	log.Printf("Successfully sent email to %s", to)
	return nil
}
