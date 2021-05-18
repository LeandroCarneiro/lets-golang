package main

import (
	domain "GoConvert/Domain"
	"fmt"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("test")
	a := domain.NewRate(1, "EUR", "BRL", 6.40, time.Now())
	b := domain.LogExecution{
		Id:   1,
		Time: time.Now(),
	}
	fmt.Println("Time: " + b.Time.String())
	fmt.Println("Currency: " + a.CurrencyFrom)
	handleRequests()
}
