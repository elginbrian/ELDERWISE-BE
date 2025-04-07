package providers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"time"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type SMTPProvider struct {
	config *config.EmailConfig
}

func NewSMTPProvider(config *config.EmailConfig) *SMTPProvider {
	return &SMTPProvider{
		config: config,
	}
}

func (p *SMTPProvider) TestConnection() error {
	addr := fmt.Sprintf("%s:%s", p.config.Host, p.config.Port)
	
	log.Printf("Testing SMTP connection to %s...", addr)
	
	// First try a simple TCP connection to see if the server is reachable
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		log.Printf("SMTP TCP connection failed: %v", err)
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()
	
	log.Printf("SMTP TCP connection successful, testing SMTP handshake...")
	
	// Now try an actual SMTP connection with an increased timeout
	client, err := smtp.Dial(addr)
	if err != nil {
		log.Printf("SMTP handshake failed: %v", err)
		return fmt.Errorf("failed to establish SMTP handshake: %w", err)
	}
	defer client.Close()
	
	log.Printf("SMTP handshake successful, connection verified")
	
	return nil
}

func (p *SMTPProvider) SendEmail(to, subject, htmlBody string) error {
	log.Printf("Sending email via SMTP to %s", to)
	
	addr := fmt.Sprintf("%s:%s", p.config.Host, p.config.Port)
	
	auth := smtp.PlainAuth("", p.config.Username, p.config.Password, p.config.Host)
	
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", p.config.FromName, p.config.FromEmail)
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"
	
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody
	
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         p.config.Host,
	}
	
	c, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("SMTP dial error: %w", err)
	}
	defer c.Close()
	
	if p.config.Secure {
		if err = c.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("SMTP TLS error: %w", err)
		}
	}
	
	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication error: %w", err)
	}
	
	if err = c.Mail(p.config.FromEmail); err != nil {
		return fmt.Errorf("SMTP sender error: %w", err)
	}
	
	if err = c.Rcpt(to); err != nil {
		return fmt.Errorf("SMTP recipient error: %w", err)
	}
	
	wc, err := c.Data()
	if err != nil {
		return fmt.Errorf("SMTP data error: %w", err)
	}
	defer wc.Close()
	
	_, err = wc.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("SMTP write error: %w", err)
	}
	
	log.Printf("Email successfully sent via SMTP to %s", to)
	return nil
}

func (p *SMTPProvider) SendEmailAsync(to, subject, htmlBody string) {
	go func() {
		if err := p.SendEmail(to, subject, htmlBody); err != nil {
			log.Printf("Error sending email asynchronously via SMTP: %v", err)
		}
	}()
}
