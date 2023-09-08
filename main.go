package main

import (
	"log"

	"github.com/halosatrio/bebop/config"
	"github.com/halosatrio/bebop/service"
	"github.com/halosatrio/bebop/service/delivery"
	postgres "github.com/halosatrio/bebop/service/repository"
)

func main() {
	// Initialize configurations
	config.InitConfig()

	// Initialize PostgreSQL connection
	db := postgres.ConnectDB()
	defer db.Close()

	// Initialize Gin router
	router := delivery.InitRouter()

	var generalUseCase service.GeneralUseCase
	delivery.CreateHandler(router, generalUseCase)

	// Run the Gin server
	router.Run()

	log.Fatal(router.Run(":8080"))
}
