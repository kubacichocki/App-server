package main

import (
	"log"
	"myrestapi/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/convert", handlers.HandleConversion)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
