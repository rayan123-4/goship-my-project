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

	// When someone visits /animals, run animalHandler.
	http.HandleFunc("/animals", animalHandler)

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

    // Check if there was an error loading the index.html page.
	if err != nil {
		// If error, send error message to browser.
		http.Error(w, "Could not load page", http.StatusInternalServerError)
	}
}

// animalHandler responds when someone visits /animals.
func animalHandler(w http.ResponseWriter, r *http.Request) {

	// Send the animal slice from animals.go to the aniaml.html page.
	err := templates.ExecuteTemplate(w, "animals.html", animals)

    // Check if there was an error loading the animals.html page.
	if err != nil {
		// If error, send error message to browser.
		http.Error(w, "Could not load animals page", http.StatusInternalServerError)
	}
}
