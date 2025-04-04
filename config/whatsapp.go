package config

import (
	"os"
)

type WhatsAppConfig struct {
	AccountSID  string
	AuthToken   string
	FromNumber  string
	ApiEndpoint string
}

func NewWhatsAppConfig() *WhatsAppConfig {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE_FROM")
	
	return &WhatsAppConfig{
		AccountSID:  accountSID,
		AuthToken:   authToken,
		FromNumber:  fromNumber,
		ApiEndpoint: "https://api.twilio.com/2010-04-01/Accounts/",
	}
}
