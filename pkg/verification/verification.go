package verification

import (
	"auth_service/config"
	"fmt"
	"net/smtp"
)

func SendVerificationToEmail(cfg *config.Config, email, body string) error {
	from := "kupalovv.muhammadjon@gmail.com"
	password := cfg.APP_PASSWORD
	subject := "Reset Password Localeats.uz"

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Set up email content.
	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	// Send the email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
