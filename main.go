package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	// for the main path we want to serve index.html
	mux.HandleFunc("/", responseFromFile("dist/index.html"))
	// Serve static files from a "static" directory

	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	component := hello("Ulf")
	mux.Handle("/home", templ.Handler(component))
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", contactHandler)

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8081"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	// Use the CORS handler
	handler := c.Handler(mux)

	err := http.ListenAndServe(":8081", handler)
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
		// Print the filePath to check if it's correct
		fmt.Println("File path:", filePath)

		// Read from file
		html, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "File not found", 404)
			return
		}
		respond(w, r, string(html))
	}
}
