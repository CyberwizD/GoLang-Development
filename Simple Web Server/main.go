package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s \n", r.URL.Path)
	fmt.Fprintf(w, "URL Path: %q", html.EscapeString(r.URL.Path))
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Web Server spinning up...")

	HandleRequests()
}
