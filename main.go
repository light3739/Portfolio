package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Read the contents of index.html and header.html from the templates folder
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/tools.html", "templates/about.html")
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
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
