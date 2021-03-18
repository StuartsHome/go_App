package main

import "net/http"

func main() {
	// var animal string
	// animal = "Lion"
	// fmt.Println(animal)
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8081", nil)
}
