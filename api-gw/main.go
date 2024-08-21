package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/pradiptarana/api-gw/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	// Register routes
	router.RegisterRoutes(r)

	// Start the server
	log.Println("API Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8088", r))
}
