package main

/*
Decorators essentially allow you to wrap existing functionality
and append or prepend your own custom functionality on top.
*/

// A Simple Decorator

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func myfunc() {
	fmt.Printf("Type %T \n", myfunc)
	time.Sleep(1 * time.Second)
}

func decoratorfunc(d func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	d()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func init() {
	fmt.Println("Decorating function...")

	decoratorfunc(myfunc)

	fmt.Println("Decorated function.")
}

func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func authPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authentication Page")
}

func isAuthorizedHeaderTrue(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking if Authorized Header is set.")

		AuthorizedHeader := r.Header.Get("Authorized")

		if AuthorizedHeader == "true" { // Correct Logic
			if value, ok := r.Header["Authorized"]; ok {
				fmt.Println(value)

				if value[0] == "true" {
					fmt.Println("Authorized Header Set")
				}

				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!")
			fmt.Fprintf(w, "Not Authorized!")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorizedHeaderTrue(Page))
	http.Handle("/auth", isAuthorizedHeaderTrue(authPage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Decorating HTTP Web Server")

	handleRequests()
}
