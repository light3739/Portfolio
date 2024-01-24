package api

import (
	"log"
	"net/smtp"
)

func sendEmail(form ContactForm, resultChan chan<- error) {
	// Use smtpConfig variables
	smtpHost := smtpConfig.Host
	smtpPort := smtpConfig.Port
	smtpEmail := smtpConfig.Email
	smtpPassword := smtpConfig.Password
	recipientEmail := smtpConfig.RecipientEmail

	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	// Define email headers and body
	to := []string{recipientEmail}
	msg := []byte("To: " + recipientEmail + "\r\n" +
		"Subject: New Contact Form Submission\r\n\r\n" +
		"From: " + form.FirstName + " " + form.LastName + " <" + form.Email + ">\r\n" +
		"Message: " + form.Message)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, to, msg)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		resultChan <- err
		return
	}

	log.Println("Email sent successfully")
	resultChan <- nil
}
