package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elginbrian/ELDERWISE-BE/config"
)

type WhatsAppService interface {
	SendMessage(to, message string) error
}

type whatsAppService struct {
	config *config.WhatsAppConfig
}

func NewWhatsAppService(config *config.WhatsAppConfig) WhatsAppService {
	return &whatsAppService{
		config: config,
	}
}

func (s *whatsAppService) SendMessage(to, message string) error {
	if !strings.HasPrefix(to, "whatsapp:") {
		to = fmt.Sprintf("whatsapp:%s", to)
	}
	
	from := s.config.FromNumber
	if !strings.HasPrefix(from, "whatsapp:") {
		from = fmt.Sprintf("whatsapp:%s", from)
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

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send WhatsApp message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("failed to decode error response, status: %d", resp.StatusCode)
		}
		return fmt.Errorf("WhatsApp API error: %v", errorResponse)
	}

	return nil
}
