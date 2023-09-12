package main

import (
	"fmt"
	"log"

	"github.com/sarweshmaharjan/api-simulator.git/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("API URL: %s\n", cfg.APIURL)
	fmt.Printf("Port: %s\n", cfg.Port)

	// Use other configuration fields as needed
}
