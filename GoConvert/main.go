package main

import (
	di "GoConvert/Infra/DI"
	serv "GoConvert/Service"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func init() {
	di.Register()
}

func main() {
	fmt.Println("test")
	result := serv.Find(1)

	fmt.Println("Time: " + result.Date.String())
	fmt.Println("Currency: " + result.CurrencyFrom)
	handleRequests()
}
