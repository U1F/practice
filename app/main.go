package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/button", buttonHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/change", changeHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	html := `<div data-function="test">
		<p data-function="change" data-event="mouseover">TEST</p>
		</div>`
	fmt.Fprint(w, html)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	html := `<p>TEST TEST TEST </p>`
	fmt.Fprint(w, html)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	html := `<p>CHANGED</p>`
	fmt.Fprint(w, html)
}
