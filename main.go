package main

import (
	"edwardlayaodev/gohtmx/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.HomeHandler())

	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	
}