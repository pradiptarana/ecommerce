package handlers

import (
	"io"
	"net/http"
	"os"
)

// User service handlers
func UserLogin(w http.ResponseWriter, r *http.Request) {
	forwardRequest(w, r, os.Getenv("USER_SERVICE_URL")+"/login")
}

// Product service handlers
func ListProducts(w http.ResponseWriter, r *http.Request) {
	forwardRequest(w, r, os.Getenv("PRODUCT_SERVICE_URL")+"/products")
}

// Order service handlers
func Checkout(w http.ResponseWriter, r *http.Request) {
	forwardRequest(w, r, os.Getenv("ORDER_SERVICE_URL")+"/checkout")
}

func ReleaseStock(w http.ResponseWriter, r *http.Request) {
	forwardRequest(w, r, os.Getenv("ORDER_SERVICE_URL")+"/release")
}

// Warehouse service handlers
func TransferProducts(w http.ResponseWriter, r *http.Request) {
	forwardRequest(w, r, os.Getenv("WAREHOUSE_SERVICE_URL")+"/transfer")
}

// Forward the request to the appropriate microservice
func forwardRequest(w http.ResponseWriter, r *http.Request, serviceURL string) {
	resp, err := http.DefaultClient.Do(r.Clone(r.Context()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
