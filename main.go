package main

import (
	"fmt"

	"log"
	"net/http"
	"os"
)

func main() {
	endpoints := map[string]string{
		"/home":    "home.html",
		"/about":   "about.html",
		"/contact": "contact.html",
	}

	for endpoint, filePath := range endpoints {
		http.HandleFunc(endpoint, generateHandler(filePath))
	}

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func generateHandler(filePath string) http.HandlerFunc {
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

func respond(w http.ResponseWriter, r *http.Request, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, html)
}
