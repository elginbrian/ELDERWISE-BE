package config

import (
	"os"
	"strings"
)

const MaxTrialMessageLength = 160

type SMSConfig struct {
	AccountSID      string
	AuthToken       string
	FromNumber      string
	ApiEndpoint     string
	TrialAccount    bool
	MaxMessageLength int
}

func NewSMSConfig() *SMSConfig {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE_FROM")
	
	trialMode := true
	if strings.ToLower(os.Getenv("TWILIO_TRIAL_ACCOUNT")) == "false" {
		trialMode = false
	}
	
	maxLength := 1600 
	if trialMode {
		maxLength = MaxTrialMessageLength
	}
	
	return &SMSConfig{
		AccountSID:      accountSID,
		AuthToken:       authToken,
		FromNumber:      fromNumber,
		ApiEndpoint:     "https://api.twilio.com/2010-04-01/Accounts/",
		TrialAccount:    trialMode,
		MaxMessageLength: maxLength,
	}
}

func (c *SMSConfig) TruncateMessage(message string) string {
	if len(message) <= c.MaxMessageLength {
		return message
	}
	
	return message[:c.MaxMessageLength-4] + "..."
}
