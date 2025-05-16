package main

import (
	"log"
	"net/http"

	"github.com/sumityadav29/taskalley/config"
)

func main() {
	server := NewServer()

	log.Printf("starting server on port %s", config.Load().Port)
	if err := http.ListenAndServe(":"+config.Load().Port, server); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
