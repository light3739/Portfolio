package api

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func sendEmail(form ContactForm, resultChan chan<- error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	recipientEmail := os.Getenv("RECIPIENT_EMAIL")

	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	// Define email headers and body
	to := []string{recipientEmail}
	msg := []byte("To: " + recipientEmail + "\r\n" +
		"Subject: New Contact Form Submission\r\n\r\n" +
		"From: " + form.FirstName + " " + form.LastName + " <" + form.Email + ">\r\n" +
		"Message: " + form.Message)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, to, msg)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		resultChan <- err
	}

	log.Println("Email sent successfully")
	resultChan <- nil
}
