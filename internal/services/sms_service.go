package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type SMSService interface {
	SendMessage(to, message string) error
}

type smsService struct {
	config *config.SMSConfig
}

func NewSMSService(config *config.SMSConfig) SMSService {
	return &smsService{
		config: config,
	}
}

func (s *smsService) SendMessage(to, message string) error {
	if !strings.HasPrefix(to, "+") {
		to = "+" + to
	}
	
	from := s.config.FromNumber
	if !strings.HasPrefix(from, "+") {
		from = "+" + from
	}

	formData := url.Values{}
	formData.Set("To", to)
	formData.Set("From", from)
	formData.Set("Body", message)

	endpoint := fmt.Sprintf("%s%s/Messages.json", s.config.ApiEndpoint, s.config.AccountSID)

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(s.config.AccountSID, s.config.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	log.Printf("Sending SMS to %s from %s", to, from)
	
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send SMS message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("failed to decode error response, status: %d", resp.StatusCode)
		}
		return fmt.Errorf("SMS API error: %v", errorResponse)
	}

	log.Printf("Successfully sent SMS to %s", to)
	return nil
}
