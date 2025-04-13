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
	
	timeout := p.config.HealthCheckTimeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}
	
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Printf("SMTP connection timed out: this usually indicates network restrictions")
			return fmt.Errorf("SMTP connection timeout (likely blocked by network/firewall): %w", err)
		}
		
		log.Printf("SMTP connection failed: %v", err)
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()
	
	log.Printf("SMTP TCP connection successful")
	return nil
}

func (p *SMTPProvider) SendEmail(to, subject, htmlBody string) error {
	log.Printf("Sending email via SMTP to %s", to)
	
	if p.config.Host == "smtp.gmail.com" {
		return p.sendGmailSimple(to, subject, htmlBody)
	}
	
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

func (p *SMTPProvider) sendGmailSimple(to, subject, htmlBody string) error {
	from := p.config.FromEmail
	password := p.config.Password
	
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	
	msg := []byte("From: " + p.config.FromName + " <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		htmlBody + "\r\n")
	
	err1 := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
	if err1 == nil {
		log.Printf("Email sent successfully via port 587")
		return nil
	}
	
	err2 := smtp.SendMail("smtp.gmail.com:465", auth, from, []string{to}, msg)
	if err2 == nil {
		log.Printf("Email sent successfully via port 465")
		return nil
	}
	
	err3 := smtp.SendMail("smtp.gmail.com:25", auth, from, []string{to}, msg)
	if err3 == nil {
		log.Printf("Email sent successfully via port 25")
		return nil
	}
	
	log.Printf("Failed to send via 587: %v", err1)
	log.Printf("Failed to send via 465: %v", err2)
	log.Printf("Failed to send via 25: %v", err3)
	
	return fmt.Errorf("all SMTP sending attempts failed")
}

func (p *SMTPProvider) SendEmailAsync(to, subject, htmlBody string) {
	go func() {
		if err := p.SendEmail(to, subject, htmlBody); err != nil {
			log.Printf("Error sending email asynchronously via SMTP: %v", err)
		}
	}()
}


