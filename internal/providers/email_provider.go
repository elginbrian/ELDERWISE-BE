package providers

type EmailProvider interface {
	SendEmail(to, subject, htmlBody string) error
	SendEmailAsync(to, subject, htmlBody string)
}

