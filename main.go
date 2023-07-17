package main

import (
	"log"
	"net/http"
	"webapp/handlers"
)


func main(){
	http.HandleFunc("/", handlers.SubscriptionHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}