package main

import (
	"log"
	"myrestapi/handlers"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Create a new router
	router := http.NewServeMux()

	// Register the handler for the /convert endpoint
	router.HandleFunc("/convert", handlers.HandleConversion)

	// Use CORS middleware
	c := cors.AllowAll()

	// Create a new CORS wrapped handler with the provided CORS options
	handler := c.Handler(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
