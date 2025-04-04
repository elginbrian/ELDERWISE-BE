package config

import (
	"os"
)

type SMSConfig struct {
	AccountSID  string
	AuthToken   string
	FromNumber  string
	ApiEndpoint string
}

func NewSMSConfig() *SMSConfig {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE_FROM")
	
	return &SMSConfig{
		AccountSID:  accountSID,
		AuthToken:   authToken,
		FromNumber:  fromNumber,
		ApiEndpoint: "https://api.twilio.com/2010-04-01/Accounts/",
	}
}
