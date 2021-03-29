package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// not sure if this struct is used???  - Remove ??
type ResponseOrder struct {
	ID    string `json: "id"`
	Name  string `json: "name"`
	Price int    `json: "price"`
	Shop  *Shop  `json: "shop"`
}

type Order struct {
	ID    string `json: "id"`
	Name  string `json: "name"`
	Price int    `json: "price"`
	Shop  *Shop  `json: "shop"`
}

type Shop struct {
	Shopname string `json: "shopname"`
	Postcode string `json: "postcode"`
}

var orders []Order

// all orders
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// single order
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	for _, i := range orders {
		if i.ID == params["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})

}

// create order
func createOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(100))
	json.NewEncoder(w).Encode(orders)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// delete order
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			fmt.Println("DELETE order: ", params["id"])
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			var order Order
			_ = json.NewDecoder(r.Body).Decode(&order)
			order.ID = params["id"]
			json.NewEncoder(w).Encode(orders)
			orders = append(orders, order)
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/orders", getOrders).Methods("GET")
	myRouter.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	myRouter.HandleFunc("/orders", createOrder).Methods("POST")
	myRouter.HandleFunc("/orders/{id}", updateOrder).Methods("PUT")
	myRouter.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

type AA []Order

func callAPI(times int) {
	for i := 1; i <= times; i++ {
		resp, err := http.Get("http://localhost:8082/orders")
		if err != nil {
			log.Fatalln(err)
		}
		data, _ := ioutil.ReadAll(resp.Body)
		var aa []interface{}
		err = json.Unmarshal(data, &aa)
		if err != nil {
			log.Fatalln(err)
		}
		for x := range aa {
			fmt.Println(aa[x])
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// ToDO
func postAPI() {

}

func callSingleAPI(api string, seconds int) {
	for i := 1; true; i++ {
		resp, err := http.Get(api)
		if err != nil {
			log.Fatalln(err)
		}
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("#####################")
		fmt.Println(string(data))
		fmt.Println("#####################")
		fmt.Println()
		time.Sleep(time.Second * time.Duration(seconds))

	}
}

func main() {
	// srv := api.NewServer()
	// http.ListenAndServe(":8080", srv)

	orders = append(orders, Order{ID: "111", Name: "Latte", Price: 2, Shop: &Shop{Shopname: "Costa", Postcode: "TW5 8NL"}})
	orders = append(orders, Order{ID: "222", Name: "Cappuccino", Price: 4, Shop: &Shop{Shopname: "Pret", Postcode: "AU1 0PT"}})

	handleRequests()
	// go callSingleAPI("http://api.open-notify.org/iss-now.json", 5)
	// callAPI(10000000)

}
