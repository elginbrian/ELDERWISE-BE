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
		// Try multiple times with increasing delays
		maxRetries := 3
		var lastErr error
		
		for i := 0; i < maxRetries; i++ {
			if err := s.SendMessage(to, subject, message); err != nil {
				log.Printf("ERROR: Email attempt %d failed: %v", i+1, err)
				lastErr = err
				// Exponential backoff
				time.Sleep(time.Duration(2*i+1) * time.Second)
				continue
			}
			
			log.Printf("SUCCESS: Async email to %s successfully sent on attempt %d", to, i+1)
			return
		}
		
		log.Printf("ERROR: All %d attempts to send email failed. Last error: %v", maxRetries, lastErr)
		
		if lastErr != nil {
			if altErr := s.sendViaSSL(to, subject, message); altErr != nil {
				log.Printf("ERROR: Alternative email method also failed: %v", altErr)
			} else {
				log.Printf("SUCCESS: Email sent via alternative method to %s", to)
			}
		}
	}()
}

func (s *emailService) SendMessage(to, subject, message string) error {
	log.Printf("ATTEMPT: Sending email to %s with subject: %s", to, subject)
	log.Printf("DEBUG: Using Gmail account: %s", s.config.Username)
	
	if s.config.Port == "465" {
		return s.sendViaSSL(to, subject, message)
	}
	
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)
	
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
	
	fullMessage := []byte(msg.String())
	
	log.Printf("DEBUG: Connecting to SMTP server at %s with 15s timeout", addr)
	
	dialer := &net.Dialer{
		Timeout: 15 * time.Second,
	}
	
	conn, err := dialer.Dial("tcp", addr)
	if err != nil {
		log.Printf("ERROR: Failed to connect to SMTP server: %v", err)
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()
	
	client, err := smtp.NewClient(conn, s.config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()
	
	tlsConfig := &tls.Config{
		ServerName: s.config.Host,
		MinVersion: tls.VersionTLS12,
	}
	
	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %w", err)
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
	
	_, err = wc.Write(fullMessage)
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
	
	log.Printf("SUCCESS: Email sent to %s", to)
	return nil
}

func (s *emailService) sendViaSSL(to, subject, message string) error {
	log.Printf("ATTEMPT: Sending email via SSL to %s", to)
	
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
	
	servername := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.config.Host,
	}
	
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Printf("ERROR: Failed to connect via SSL: %v", err)
		return err
	}
	defer conn.Close()
	
	client, err := smtp.NewClient(conn, s.config.Host)
	if err != nil {
		return err
	}
	defer client.Close()
	
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)
	if err = client.Auth(auth); err != nil {
		return err
	}
	
	if err = client.Mail(s.config.FromEmail); err != nil {
		return err
	}
	
	if err = client.Rcpt(to); err != nil {
		return err
	}
	
	w, err := client.Data()
	if err != nil {
		return err
	}
	
	_, err = w.Write([]byte(msg.String()))
	if err != nil {
		return err
	}
	
	err = w.Close()
	if err != nil {
		return err
	}
	
	client.Quit()
	
	log.Printf("SUCCESS: Email sent via SSL to %s", to)
	return nil
}

// Add a new helper method to determine if error is a network error
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	
	_, isNetErr := err.(net.Error)
	return isNetErr || 
		strings.Contains(err.Error(), "dial tcp") || 
		strings.Contains(err.Error(), "connection") ||
		strings.Contains(err.Error(), "timeout")
}
