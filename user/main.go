package main

import (
	"log"

	"github.com/pradiptarana/user/app"
)

func main() {
	if err := app.SetupServer().Run("localhost:8081"); err != nil {
		log.Fatal(err)
	}
}
