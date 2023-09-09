package main

import (
	"github.com/halosatrio/bebop/config"
	"github.com/halosatrio/bebop/router"
)

func main() {
	config.LoadEnv()
	r := router.SetupRouter()

	db := config.ConnectDB()
	defer db.Close()

	r.Run(":8080")
}
