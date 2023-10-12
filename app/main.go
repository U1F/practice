package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/button", buttonHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Add this line

	html := `<div>TEST</div>`
	fmt.Fprint(w, html)
}
