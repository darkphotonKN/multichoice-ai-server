package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkphotonKN/go-ollama-chat/config"
)

func main() {

	// setup config
	cfg := config.LoadConfig()

	// load routers
	routeHandler := SetupRoutes(&cfg)

	log.Printf("Starting on port %s\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), routeHandler)
}
