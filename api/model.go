package api

type ContactForm struct {
	FirstName string
	LastName  string
	Email     string
	Message   string
}

type SMTPConfig struct {
	Host           string
	Port           string
	Email          string
	Password       string
	RecipientEmail string
}
