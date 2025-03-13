package main

/*
JWTs, or JSON Web Tokens are a compact, URL-safe means of representing claims
which allows the transmition of information from a client to the server in a stateless, but secure way.

The JWT standard uses either a secret, using the HMAC algorithm, or a public/private key pair using RSA or ECDSA.

JWTs are heavily used within Single-Page Applications (SPAs) as a means of secure communications as they allow two key things:

1). Authentication - The most commonly used practice. Once a user logs into various applications,
	every request that is then sent from the client on behalf of the user will contain the JWT.
2). Information Exchange - The second use for JWTs is to securely transmit information between different systems.
	These JWTs can be signed using public/private key pairs so you can verify each system in this transaction in a secure manner
	and JWTs contain an anti-tamper mechanism as they are signed based off the header and the payload.
*/

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt"
)

var SigningKey = []byte("randomsecretkey")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World \n")
	fmt.Printf("Endpoint: %s \n", r.URL.Path)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Token")
		// This code checks if the "Token" header is present and not empty before attempting to parse it.

		if tokenHeader != "" {
			token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}

				return SigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, "%s", err.Error())
				return
			}

			if token.Valid {
				endpoint(w, r)
			} else {
				fmt.Fprintf(w, "Not Authorized")
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Server starting..")

	handleRequests()
}
