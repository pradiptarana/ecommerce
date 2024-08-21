package router

import (
	"github.com/gorilla/mux"
	"github.com/pradiptarana/api-gw/handlers"
)

func RegisterRoutes(r *mux.Router) {
	// User service routes
	r.HandleFunc("/login", handlers.UserLogin).Methods("POST")

	// Product service routes
	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")

	// Order service routes
	r.HandleFunc("/checkout", handlers.Checkout).Methods("POST")

	// Warehouse service routes
	r.HandleFunc("/transfer", handlers.TransferProducts).Methods("POST")
}
