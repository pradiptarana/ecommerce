package main

import (
	"log"

	"github.com/pradiptarana/order/app"
)
	
func main() {
	if err := app.SetupServer().Run("localhost:8083"); err != nil {
		log.Fatal(err)
	}
}
