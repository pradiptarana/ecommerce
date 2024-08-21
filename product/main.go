package main

import (
	"log"

	"github.com/pradiptarana/product/app"
)

func main() {
	if err := app.SetupServer().Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
