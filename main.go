package main

import (
	"html/template"
	"log"
	"net/http"
)


func main() {
	// When someone visits / run homeHandler
	http.HandleFunc("/", animalHandler)

	http.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),

	))

	createBarChart()

        http.HandleFunc("/bar.html", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "bar.html")
    })

	log.Println("BioGO is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// animalHandler responds when someone visits animal.html.
func animalHandler(w http.ResponseWriter, r *http.Request) {

    // Load html templates on page
    templates, err := template.ParseGlob("templates/*.html")

    // Check if there was an error loading the animals.html page.
	if err != nil {
		// If error, send error message to browser.
		http.Error(w, "Could not load animals page", http.StatusInternalServerError)
	}

	// Send the animal slice from animals.go to the animal.html page.
	err = templates.ExecuteTemplate(w, "animal.html", animals)

	// Check if there was an error loading the animals.html page.
	if err != nil {
		// If error, send error message to browser.
		http.Error(w, "Could not load animals page", http.StatusInternalServerError)
	}

}
