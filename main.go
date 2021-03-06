package main

import (
	"log"

	"github.com/christianXpinoza/simple-card-demo/api"
	"github.com/christianXpinoza/simple-card-demo/card"
)

func main() {
	// Create a new card service
	cardService := card.New()
	// New webApi service instance
	webAPI := api.New()
	// Start Web API
	if err := webAPI.Start(":8080", cardService); err != nil {
		log.Fatal(err)
	}

}
