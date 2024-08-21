package main

import (
	"log"

	"github.com/pradiptarana/warehouse/app"
)

func main() {
	if err := app.SetupServer().Run("localhost:8084"); err != nil {
		log.Fatal(err)
	}
}
