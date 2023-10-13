package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/a-h/templ"
)

func main() {
	// for the main path we want to serve index.html
	http.HandleFunc("/", responseFromFile("index.html"))
// Serve static files from a "static" directory
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

	component := hello("Ulf")
	http.Handle("/home", templ.Handler(component))
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



func aboutHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, r, "<h1>About Us</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, r, "<h1>Contact</h1>")
}

func respond(w http.ResponseWriter, r *http.Request, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, html)
}

func responseFromFile(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Read from file
		html, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		respond(w, r, string(html))
	}
}