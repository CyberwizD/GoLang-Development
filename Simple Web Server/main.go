package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	counter int
	mutex   = &sync.Mutex{}
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, r.URL.Path[2:])

	fmt.Fprintf(w, "URL: %s \n", r.URL.Path)
	fmt.Fprintf(w, "URL Path: %q", html.EscapeString(r.URL.Path))
}

func CounterPage(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	counter++

	fmt.Fprintf(w, "You have visited this page %s times.", strconv.Itoa(counter))

	mutex.Unlock()
}

func HandleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/Home", HomePage)

	http.HandleFunc("/Counter", CounterPage)

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)) // Adding a secure connection - HTTPS

	// Using openssl to generate self-signed certs
	// $ openssl genrsa -out server.key 2048
	// $ openssl ecparam -genkey -name secp384r1 -out server.key
	// $ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
}

func main() {
	fmt.Println("Web Server spinning up...")

	HandleRequests()
}
