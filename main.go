package main

import (
	"Portfolio/api"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Read the contents of index.html and header.html from the templates folder
	tmpl, err := template.ParseFiles(
		"templates/index.html", "templates/header.html",
		"templates/tools.html", "templates/about.html",
		"templates/calculator.html",
		"templates/contact.html",
		"templates/calculator_popup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	//apiKey := "your_api_key_here" // Replace with your actual API key TODO: fix AuthMiddleware
	contactHandler := http.HandlerFunc(api.ContactFormHandler)
	contactHandlerWithMiddleware := api.LoggingMiddleware(contactHandler)

	http.Handle("/contact", contactHandlerWithMiddleware)
	http.HandleFunc("/calculator-popup", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/calculator_popup.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
