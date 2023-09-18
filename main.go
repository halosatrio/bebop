package main

import (
	"github.com/halosatrio/bebop/config"
	"github.com/halosatrio/bebop/handlers"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDB()
	defer db.Close()

	r := handlers.SetupRouter(db)
	r.Run(":8080")
}
