package email

import (
	"fmt"
	"net/smtp"
	"github.com/MaryJane-09/nexus/backend/internal/otp"
)

type EmailSender struct {
	GmailAddress     string
	GmailAppPassword string
	SMTPHost         string
	SMTPPort         string
}

func (s EmailSender) SendVerification(email string, code string) error {
	auth := smtp.PlainAuth("", s.GmailAddress, s.GmailAppPassword, s.SMTPHost)
	headerFrom := fmt.Sprintf("From: %s\r\n", s.GmailAddress)
	headerTo := fmt.Sprintf("To: %s\r\n", email)
	headerSubject := "Subject: Your Verification Code\r\n"
	headerMime := "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n"

	body := fmt.Sprintf("\nHello,\n\nYour one-time verification code is: %s\n\nThis code will expire by %d. Do not share this code  with anyone.\n\nif you did not request for this please ignore", code, otp.ExpiryTime)
	messageBytes := []byte(headerFrom + headerTo + headerSubject + headerMime + body)

	err := smtp.SendMail(s.SMTPHost+":"+
		s.SMTPPort, auth, s.GmailAddress,
		[]string{email}, messageBytes)
	if err != nil {
		return fmt.Errorf("failed to send native email: %w", err)
	}

	return nil
}
