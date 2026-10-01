package main

import (
	"html/template"
	"log"
	"net/http"
)

// Look in templates and find html files
// Send html to browser
var templates = template.Must(template.ParseGlob("templates/*.html"))


func main() {
	// When someone visits / run homeHandler
	http.HandleFunc("/", homeHandler)

	http.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),

	))

	log.Println("BioGO is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

 // Func that responds to the visitor
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
	}
}

