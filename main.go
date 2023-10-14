package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/rs/cors"
	"grapefruixyz.org/m/v2/app"
	"grapefruixyz.org/m/v2/config"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Use the configuration
	fmt.Println("Server Address:", cfg.ServerAddress)

	// Create the HTTP server
	mux := http.NewServeMux()

	// for the main path we want to serve index.html
	mux.HandleFunc("/", responseFromFile("dist/index.html"))
	// Serve static files from a "static" directory

	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

	home := app.Home("Ulf")
	mux.Handle("/home", templ.Handler(home))

	about := app.About("about")
	mux.Handle("/about", templ.Handler(about))

	contact := app.Contact("contact")
	mux.Handle("/contact", templ.Handler(contact))

	// CORS setup using values from the configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: cfg.AllowCredentials,
		AllowedMethods:   cfg.AllowedMethods,
	})

	// Use the CORS handler
	handler := c.Handler(mux)

	err = http.ListenAndServe(":8081", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func respond(w http.ResponseWriter, r *http.Request, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, html)
}

func responseFromFile(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		html, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "File not found", 404)
			return
		}
		respond(w, r, string(html))
	}
}
