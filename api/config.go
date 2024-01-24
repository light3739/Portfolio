package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var smtpConfig SMTPConfig

func LoadConfig() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	smtpConfig = SMTPConfig{
		Host:           os.Getenv("SMTP_HOST"),
		Port:           os.Getenv("SMTP_PORT"),
		Email:          os.Getenv("SMTP_EMAIL"),
		Password:       os.Getenv("SMTP_PASSWORD"),
		RecipientEmail: os.Getenv("RECIPIENT_EMAIL"),
	}
}
