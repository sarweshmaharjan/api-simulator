package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sarweshmaharjan/api-simulator.git/api"
	"github.com/sarweshmaharjan/api-simulator.git/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("API URL: %s\n", cfg.APIURL)
	fmt.Printf("Port: %s\n", cfg.Port)

	router := gin.Default()
	api.Routes(router)
	uri := fmt.Sprintf("%s:%s", cfg.APIURL, cfg.Port)
	router.Run(uri)
}
