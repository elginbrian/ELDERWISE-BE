package providers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type MailgunProvider struct {
	apiKey      string
	domain      string
	fromEmail   string
	fromName    string
}

func NewMailgunProvider(config *config.EmailConfig) *MailgunProvider {
	return &MailgunProvider{
		apiKey:    config.MailgunAPIKey,
		domain:    config.MailgunDomain,
		fromEmail: config.FromEmail,
		fromName:  config.FromName,
	}
}

func (p *MailgunProvider) SendEmail(to, subject, htmlBody string) error {
	if p.apiKey == "" || p.domain == "" {
		return fmt.Errorf("mailgun API key or domain is not configured")
	}

	apiURL := fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", p.domain)
	
	form := url.Values{}
	form.Add("from", fmt.Sprintf("%s <%s>", p.fromName, p.fromEmail))
	form.Add("to", to)
	form.Add("subject", subject)
	form.Add("html", htmlBody)
	
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth("api", p.apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 400 {
		return fmt.Errorf("mailgun API returned error status: %d", resp.StatusCode)
	}
	
	return nil
}

func (p *MailgunProvider) SendEmailAsync(to, subject, htmlBody string) {
	go func() {
		if err := p.SendEmail(to, subject, htmlBody); err != nil {
			log.Printf("Error sending email via Mailgun asynchronously: %v", err)
		}
	}()
}
