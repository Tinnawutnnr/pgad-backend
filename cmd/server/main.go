package main

import (
	"log"

	"pgad/internal/app"
)

func main() {
	a := app.New()
	log.Println("listening on :8080")
	if err := a.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
