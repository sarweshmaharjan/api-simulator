package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sarweshmaharjan/api-simulator.git/internal/api"
	"github.com/sarweshmaharjan/api-simulator.git/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("API URL: %s\n", cfg.APIURL)
	fmt.Printf("Port: %s\n", cfg.Port)

	router := gin.Default()
	router.POST("/v1/direct_transfer", api.GetDirectTransfer)
	uri := fmt.Sprintf("%s:%s", cfg.APIURL, cfg.Port)
	router.Run(uri)
}
