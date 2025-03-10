package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desp"`
	Content     string `json:"content"`
}

var Articles []Article

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint: HomePage")
}

func ReturnArticlePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: ArticlePage")
	json.NewEncoder(w).Encode(Articles)
}

func GetArticleIDPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enpoint: ArrticleIDPage")

	pathVar := mux.Vars(r)
	key := pathVar["id"]

	for _, articleContent := range Articles {
		if articleContent.Id == key {
			json.NewEncoder(w).Encode(articleContent)
		}
	}
}

func CreateNewArticlePage(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "%+v", string(requestBody))
}

func UpdateArticlePage(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var updateArticle Article

	json.Unmarshal(reqBody, &updateArticle)

	Articles = append(Articles, updateArticle)

	json.NewEncoder(w).Encode(Articles)
}

func RequestHandler() {
	router := mux.NewRouter()

	// http.HandleFunc("/", HomePage)
	// http.HandleFunc("/Articles", ArticlePage)

	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/Articles", ReturnArticlePage).Methods("GET")
	router.HandleFunc("/Articles", CreateNewArticlePage).Methods("POST")
	router.HandleFunc("/Articles", UpdateArticlePage)
	router.HandleFunc("/Articles/{id}", GetArticleIDPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		{
			Id:          "1",
			Title:       "GoLang",
			Description: "Simple REST API",
			Content:     "Building a simple HTTP server with the implementation of REST API for CRUD opearatoins.",
		},

		{
			Id:          "2",
			Title:       "Rust",
			Description: "Rust VS GoLang, Which is better?",
			Content:     "Rust is suitable for memory management, GoLang on the other hand is suitable for concurrency.",
		},
	}

	RequestHandler()
}
