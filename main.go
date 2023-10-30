package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sarweshmaharjan/api-simulator.git/api"
	"github.com/sarweshmaharjan/api-simulator.git/config"
	"github.com/sarweshmaharjan/api-simulator.git/storage"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Setup database connection
	storage.OpenDBConnection()

	fmt.Printf("API URL: %s\n", config.Cfg.APIURL)
	fmt.Printf("Port: %s\n", config.Cfg.Port)

	router := gin.Default()
	api.Routes(router)
	uri := fmt.Sprintf("%s:%s", config.Cfg.APIURL, config.Cfg.Port)
	router.Run(uri)
}
