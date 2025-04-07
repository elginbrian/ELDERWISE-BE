package providers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type SendGridProvider struct {
	apiKey    string
	fromEmail string
	fromName  string
}

func NewSendGridProvider(config *config.EmailConfig) *SendGridProvider {
	return &SendGridProvider{
		apiKey:    config.SendGridAPIKey,
		fromEmail: config.FromEmail,  // This can be your Gmail address
		fromName:  config.FromName,
	}
}

func (p *SendGridProvider) SendEmail(to, subject, htmlBody string) error {
	if p.apiKey == "" {
		return fmt.Errorf("SendGrid API key is not configured")
	}

	data := fmt.Sprintf(`{
		"personalizations": [{
			"to": [{"email": "%s"}],
			"subject": "%s"
		}],
		"from": {"email": "%s", "name": "%s"},
		"content": [{"type": "text/html", "value": %s}]
	}`, to, subject, p.fromEmail, p.fromName, quoteString(htmlBody))

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("SendGrid API returned error status: %d", resp.StatusCode)
	}

	return nil
}

func (p *SendGridProvider) SendEmailAsync(to, subject, htmlBody string) {
	go func() {
		if err := p.SendEmail(to, subject, htmlBody); err != nil {
			log.Printf("Error sending email via SendGrid asynchronously: %v", err)
		}
	}()
}

func quoteString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return "\"" + s + "\""
}
