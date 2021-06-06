package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string      `json:"Id"`
	Title   string      `json:"Title"`
	Desc    string      `json:"desc"`
	Content string      `json:"content"`
	Codes   []Paragraph `json: "codes"`
}

type Paragraph struct {
	Name   string    `json:"Name"`
	Prices []float64 `json:"Prices"`
}

// global Articles array
// populate in main function to simulate database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // encodes the Articles array into a JSON string and then writes as part of response
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// loop over all Articles
	// if article.Id equals the key, return article encoded as JSON

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
	fmt.Fprintf(w, "Key: "+key)
}
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":1000", myRouter))
}

func main() {

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content", Codes: []Paragraph{
			Paragraph{
				Name:   "Trevor",
				Prices: []float64{15000.00, 10000.00},
			},
			Paragraph{
				Name:   "Adam",
				Prices: []float64{20000.00},
			},
		}},
		// Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content", Codes: []string{"CC", "DD"}},
	}

	handleRequests()
}
