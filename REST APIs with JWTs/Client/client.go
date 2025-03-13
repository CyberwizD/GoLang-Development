package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var UserSigningKey = []byte("randomsecretkey")

func loginPage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://localhost:8080", nil)

	request.Header.Set("Token", validToken)
	response, err := client.Do(request)

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "UserName"
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	tokenString, err := token.SignedString(UserSigningKey)

	if err != nil {
		log.Printf("Error: Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleUserRequests() {
	http.HandleFunc("/", loginPage)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func main() {
	fmt.Println("Client call..")

	handleUserRequests()
}
