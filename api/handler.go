package api

import (
	"net/http"
)

func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form values
		form := ContactForm{
			FirstName: r.FormValue("firstName"),
			LastName:  r.FormValue("lastName"),
			Email:     r.FormValue("email"),
			Message:   r.FormValue("message"),
		}

		if form.Email == "" || form.Message == "" {
			http.Error(w, "Email and message are required", http.StatusBadRequest)
			return
		}
		resultChan := make(chan error)
		go sendEmail(form, resultChan)

		err := <-resultChan
		if err != nil {
			http.Error(w, "Error sending email", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("Form submitted successfully")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		// Handle error: only POST method is allowed
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
}
