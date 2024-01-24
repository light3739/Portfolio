package main

import (
	"Portfolio/api"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(
		"templates/index.html", "templates/header.html",
		"templates/tools.html", "templates/about.html",
		"templates/calculator.html",
		"templates/contact.html",
		"templates/calculator_popup.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	api.LoadConfig()
	contactHandler := http.HandlerFunc(api.ContactFormHandler)
	contactHandlerWithMiddleware := api.ErrorHandlingMiddleware(api.LoggingMiddleware(contactHandler))

	http.Handle("/contact", contactHandlerWithMiddleware)
	http.HandleFunc("/calculator-popup", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/calculator_popup.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server := &http.Server{
		Addr: ":8080",
		// Set timeouts to avoid Slowlori	s attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      nil, // use http.DefaultServeMux
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
