package main

import (
	"net/http"

	"github.com/gowebexamples/http-server/api"
)

func main() {
	// var animal string
	// animal = "Lion"
	// fmt.Println(animal)
	// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })
	// http.ListenAndServe(":8081", nil)

	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
