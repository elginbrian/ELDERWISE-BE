package config

import "os"

type SMSConfig struct {
	AccountSID  string
	AuthToken   string
	FromNumber  string
	ApiEndpoint string
}

func NewSMSConfig() *SMSConfig {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	if accountSID == "" {
		accountSID = "ACd0b82f570065698d03ff1ea74b24db64"
	}
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if authToken == "" {
		authToken = "a2f4884d8f19869cd869e03a07cf8e4e"
	}
	fromNumber := os.Getenv("TWILIO_SMS_FROM")
	if fromNumber == "" {
		fromNumber = "+15077044325"
	}
	return &SMSConfig{
		AccountSID:  accountSID,
		AuthToken:   authToken,
		FromNumber:  fromNumber,
		ApiEndpoint: "https://api.twilio.com/2010-04-01/Accounts/",
	}
}
