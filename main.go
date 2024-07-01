package main

import (
	"edwardlayaodev/gohtmx/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	// Load env vars 
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	// URL Routes
	http.HandleFunc("/", handlers.HomeHandler())

	// Server Start
	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	
}